package captiveportal

import (
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

/* Send and/or clear the Auth List when requested by openNDS
When a client was verified, their parameters were added to the "auth list"
The auth list is sent to openNDS when authmon requests it.

auth_get:
auth_get is sent by authmon or libopennds in a POST request and can have the following values:

1. Value "list".
	FAS sends the auth list and deletes each client entry currently on that list.

2. Value "view".
	FAS checks the received payload for an ack list of successfully authenticated clients from previous auth lists.
	Clients on the auth list are only deleted if they are in a received ack list.
	Authmon will have sent the ack list as acknowledgement of all clients that were successfully authenticated in the previous auth list.
	Finally FAS replies by sending the next auth list.
	"view" is the default method used by authmon.

3. Value "clear".
	This is a housekeeping function and is called by authmon on startup of openNDS.
	The auth list is cleared as any entries held by this FAS at the time of openNDS startup will be stale.

4. Value "deauthed".
	FAS receives a payload containing notification of deauthentication of a client and the reason for that notification.
	FAS replies with an ack., confirming reception of the notification.

5. Value "custom".
	FAS receives a payload containing a b64 encoded string to be used by FAS to provide custom functionality.
	FAS replies with an ack., confirming reception of the custom string.
*/

type AuthGetRequest struct {
	AuthGet     string `form:"auth_get"`
	GatewayHash string `form:"gatewayhash"`
	Payload     string `form:"payload"`
}

func listAuthClients(gwhash string) (string, error) {
	absBase := "/Users/henry/projects/captive-portal/authfiles"
	absBase = filepath.Join(absBase, gwhash)

	dir, err := os.ReadDir(absBase)
	if err != nil {
		return "", errors.WithStack(err)
	}

	if len(dir) == 0 {
		return "", nil
	}

	clients := make([]string, 0, len(dir))
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		body, err := os.ReadFile(filepath.Join(absBase, entry.Name()))
		if err != nil {
			return "", errors.WithStack(err)
		}

		clients = append(clients, strings.Trim(string(body), ""))
	}

	return " " + strings.Join(clients, " "), nil
}

func delAuthClient(gwhash, rhid string) error {
	absBase := "/Users/henry/projects/captive-portal/authfiles"
	p := filepath.Join(absBase, gwhash, rhid)

	return errors.WithStack(os.Remove(p))
}

const (
	AuthGetHandlerURL = "/fas-aes-https.php"
)

func NewAuthGetHandler(rCli *redis.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		var r AuthGetRequest
		if err := c.Bind(&r); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", errors.WithStack(err)))
		}

		decoded, err := base64.StdEncoding.DecodeString(r.Payload)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", errors.WithStack(err)))
		}

		r.Payload = string(decoded)

		if r.AuthGet == "view" {
			if r.Payload == "none" {
				// Listing all auth clients
				rhids, err := rCli.LRange(c.Request().Context(), r.GatewayHash, 0, -1).Result()
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", errors.WithStack(err)))
				}

				if err := rCli.Del(c.Request().Context(), r.GatewayHash).Err(); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", errors.WithStack(err)))
				}

				clients := make([]string, len(rhids))
				for i, rhid := range rhids {
					result, err := rCli.Get(c.Request().Context(), fmt.Sprintf("%s/%s", r.GatewayHash, rhid)).Result()
					if err != nil {
						return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", errors.WithStack(err)))
					}

					clients[i] = strings.Trim(result, "")
				}

				base := "*"
				if len(clients) > 0 {
					base += " "
				}

				return c.String(http.StatusOK, base+strings.Join(clients, "\n"))
			} else {
				for _, hid := range strings.Split(r.Payload, "\n") {
					hid = strings.TrimLeft(hid, "* ")
					if hid == "" {
						continue
					}

					// Use GetDel for ensuring the client is in the auth list
					if err := rCli.GetDel(c.Request().Context(), fmt.Sprintf("%s/%s", r.GatewayHash, hid)).Err(); err != nil {
						return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", errors.WithStack(err)))
					}
				}

				return c.String(http.StatusOK, "ack")
			}
		} else if r.AuthGet == "clear" {
			rhids, err := rCli.LRange(c.Request().Context(), r.GatewayHash, 0, -1).Result()
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", errors.WithStack(err)))
			}

			if err := rCli.Del(c.Request().Context(), r.GatewayHash).Err(); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", errors.WithStack(err)))
			}

			for _, rhid := range rhids {
				if err := rCli.Del(c.Request().Context(), fmt.Sprintf("%s/%s", r.GatewayHash, rhid)).Err(); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", errors.WithStack(err)))
				}
			}
		}

		c.Logger().Errorf("unexpected request: %+v", r)
		return c.JSON(http.StatusOK, r)
	}
}

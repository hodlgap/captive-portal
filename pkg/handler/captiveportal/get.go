package captiveportal

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/hodlgap/captive-portal/pkg/auth"
)

/* Send and/or clear the Auth List when requested by openNDS
When a client was verified, their parameters were added to the "auth list"
The auth list is sent to openNDS when auth-mon requests it.

auth_get:
auth_get is sent by auth-mon or lib-opennds in a POST request and can have the following values:

1. Value "list".
	FAS sends the auth list and deletes each client entry currently on that list.

2. Value "view".
	FAS checks the received payload for an ack list of successfully authenticated clients from previous auth lists.
	Clients on the auth list are only deleted if they are in a received ack list.
	Auth-mon will have sent the ack list as acknowledgement of all clients that were successfully authenticated in the previous auth list.
	Finally, FAS replies by sending the next auth list.
	"view" is the default method used by auth-mon.

3. Value "clear".
	This is a housekeeping function and is called by auth-mon on startup of openNDS.
	The auth list is cleared as any entries held by these FAS at the time of openNDS startup will be stale.

4. Value "deauthed".
	FAS receives a payload containing notification of de-authentication of a client and the reason for that notification.
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

const (
	AuthGetHandlerURL = "/fas-aes-https.php"
)

func NewAuthGetHandler(authProvider auth.Provider) echo.HandlerFunc {
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
				rhids, err := authProvider.PopAllClients(c.Request().Context(), r.GatewayHash)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
				}

				// Early return if no clients
				if len(rhids) == 0 {
					return c.String(http.StatusOK, "*")
				}

				policies, err := authProvider.ListPolicies(c.Request().Context(), r.GatewayHash, rhids...)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
				}

				for i, policy := range policies {
					policies[i] = strings.Trim(policy, "")
				}

				return c.String(http.StatusOK, "* "+strings.Join(policies, "\n"))
			} else {
				if err := authProvider.DeletePolicies(c.Request().Context(), r.GatewayHash, decodeHIDs(r.Payload)...); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
				}

				return c.String(http.StatusOK, "ack")
			}
		} else if r.AuthGet == "clear" {
			// Listing all auth clients
			rhids, err := authProvider.PopAllClients(c.Request().Context(), r.GatewayHash)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
			}

			// Early return if no clients
			if len(rhids) == 0 {
				return c.String(http.StatusOK, "*")
			}

			if err := authProvider.DeletePolicies(c.Request().Context(), r.GatewayHash, rhids...); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
			}

			return c.NoContent(http.StatusOK)
		}

		// TODO(@Kcrong): Implement other auth_get methods
		msg := fmt.Sprintf("unexpected request: %+v", r)
		c.Logger().Error(msg)
		return echo.NewHTTPError(http.StatusInternalServerError, msg)
	}
}

func decodeHIDs(payload string) []string {
	rawHIDs := strings.Split(payload, "\n")

	HIDs := make([]string, 0, len(rawHIDs))
	for _, hid := range rawHIDs {
		hid = strings.TrimLeft(hid, "* ")
		if hid == "" {
			continue
		}

		HIDs = append(HIDs, hid)
	}

	return HIDs
}

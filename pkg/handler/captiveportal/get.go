package captiveportal

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
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

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	return url.QueryEscape(string(bs[:]))
}

func NewAuthGetHandler() echo.HandlerFunc {
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

		return c.JSON(http.StatusOK, r)
	}
}

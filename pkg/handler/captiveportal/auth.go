package captiveportal

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
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

type AuthRequest struct {
	FAS string `query:"fas"`
	IV  string `query:"iv"`
}

type DecodedRequest struct {
	Authdir        string `json:"authdir"`
	ClientType     string `json:"client_type"`
	Clientif       string `json:"clientif"`
	Clientip       string `json:"clientip"`
	Clientmac      string `json:"clientmac"`
	Gatewayaddress string `json:"gatewayaddress"`
	Gatewaymac     string `json:"gatewaymac"`
	Gatewayname    string `json:"gatewayname"`
	Gatewayurl     string `json:"gatewayurl"`
	Hid            string `json:"hid"`
	Originurl      string `json:"originurl"`
	Themespec      string `json:"themespec"`
	Version        string `json:"version"`
}

func (r *DecodedRequest) LoadFromMap(m map[string]string) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, r)
	}
	return errors.WithStack(err)
}

func AES256Decode(cipherText string, encKey string, iv string) (string, error) {
	decodedText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", errors.WithStack(err)
	}
	_, err = base64.StdEncoding.Decode(decodedText, decodedText)
	if err != nil {
		return "", errors.WithStack(err)
	}

	block, err := aes.NewCipher([]byte(encKey))
	if err != nil {
		return "", errors.WithStack(err)
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(decodedText, decodedText)

	return string(decodedText[:]), nil
}

func NewAuthHandler(encryptionKey string) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := new(AuthRequest)
		if err := c.Bind(&r); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", errors.WithStack(err)))
		}

		d, err := AES256Decode(r.FAS, encryptionKey, r.IV)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
		}

		res := make(map[string]string)
		items := strings.Split(d, ", ")
		for _, item := range items {
			parts := strings.Split(item, "=")
			if len(parts) != 2 {
				continue
			}
			res[parts[0]] = parts[1]
		}

		dr := DecodedRequest{}
		if err := dr.LoadFromMap(res); err != nil {
			return err
		}

		return c.JSON(200, dr)
	}
}

package captiveportal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	echo "github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	redis "github.com/redis/go-redis/v9"
)

type RawAuthRequest struct {
	FAS string `query:"fas"`
	IV  string `query:"iv"`
}

type DecodedAuthRequest struct {
	AuthDir         string `json:"authdir"`
	ClientType      string `json:"client_type"`
	ClientInterface string `json:"clientif"`
	ClientIP        string `json:"clientip"`
	ClientMAC       string `json:"clientmac"`
	GatewayAddress  string `json:"gatewayaddress"`
	GatewayMAC      string `json:"gatewaymac"`
	GatewayName     string `json:"gatewayname"`
	GatewayURL      string `json:"gatewayurl"`
	HID             string `json:"hid"`
	OriginURL       string `json:"originurl"`
	ThemeSpec       string `json:"themespec"`
	Version         string `json:"version"`
}

func (dr *DecodedAuthRequest) FromEchoContext(c echo.Context, key string) error {
	r := new(RawAuthRequest)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", errors.WithStack(err)))
	}

	d, err := AES256Decode(r.FAS, key, r.IV)
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

	data, err := json.Marshal(res)
	if err == nil {
		err = json.Unmarshal(data, dr)
	}
	return errors.WithStack(err)
}

func AES256Decode(cipherText string, encKey string, iv string) (string, error) {
	decodedText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if _, err := base64.StdEncoding.Decode(decodedText, decodedText); err != nil {
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

func hash(s string) string {
	bs := sha256.Sum256([]byte(s))

	return hex.EncodeToString(bs[:])
}

const (
	AuthHandlerURL = "/fas-aes-https.php"
)

func NewAuthHandler(encryptionKey string, rCli *redis.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		dr := new(DecodedAuthRequest)
		if err := dr.FromEchoContext(c, encryptionKey); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
		}

		rhid := hash(strings.Trim(dr.HID, "") + strings.Trim(encryptionKey, ""))

		sessionlength := 1001 // 1001 minutes
		uploadrate := 1002    // Upload Rate Limit Threshold: 1002 Kb/s
		downloadrate := 1003  // Download Rate Limit Threshold: 1003 Kb/s
		uploadquota := 1004   // Upload Quota: 1004 KBytes
		downloadquota := 1005 // Download Quota: 1005 KBytes
		custom := "key=value, key2=value2"
		custom = base64.StdEncoding.EncodeToString([]byte(custom))

		returnStr := fmt.Sprintf(
			"%s %d %d %d %d %d %s",
			rhid,
			sessionlength,
			uploadrate,
			downloadrate,
			uploadquota,
			downloadquota,
			custom,
		)

		// url encode
		returnStr = url.QueryEscape(returnStr)
		returnStr = strings.ReplaceAll(returnStr, "+", "%20")

		gwHash := hash(dr.GatewayName)

		if err := rCli.LPush(c.Request().Context(), gwHash, rhid).Err(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
		}
		key := fmt.Sprintf("%s/%s", gwHash, rhid)
		if err := rCli.Set(c.Request().Context(), key, returnStr, 0).Err(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
		}

		return c.JSON(200, dr)
	}
}

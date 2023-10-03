package captiveportal

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/hodlgap/captive-portal/pkg/auth"
	"github.com/hodlgap/captive-portal/pkg/models"
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

// MustGetAuthorizedPath returns the path that client redirects after authentication
func (dr *DecodedAuthRequest) MustGetAuthorizedPath() string {
	path, err := url.JoinPath(dr.GatewayURL, dr.AuthDir)
	if err != nil {
		log.Errorf("%+v", errors.Wrapf(err, "failed to join path. gatewayURL: %s, authDir: %s", dr.GatewayURL, dr.AuthDir))
		return dr.GatewayURL
	}
	return path
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

		unescaped, err := url.QueryUnescape(parts[1])
		if err != nil {
			unescaped = parts[1]
		}

		res[parts[0]] = unescaped
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

func NewAuthHandler(encryptionKey string, authProvider auth.Provider, db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		dr := new(DecodedAuthRequest)
		if err := dr.FromEchoContext(c, encryptionKey); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
		}

		rhid := hash(strings.Trim(dr.HID, "") + strings.Trim(encryptionKey, ""))

		p := auth.NewClientPolicy(
			rhid,
			10,
			5000,
			5000,
			5000,
			5000,
			map[string]any{
				"now": time.Now().UTC().Format(time.RFC3339),
			},
		)

		if err := authProvider.AddPolicy(c.Request().Context(), dr.GatewayName, rhid, p); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
		}

		if err := insertAuthAttemptLog(c.Request().Context(), db, *dr, *p); err != nil {
			c.Logger().Errorf("%+v", errors.WithStack(err))
		}

		return c.Redirect(http.StatusFound, dr.MustGetAuthorizedPath())
	}
}

func insertAuthAttemptLog(ctx context.Context, db *sql.DB, dr DecodedAuthRequest, p auth.ClientPolicy) error {
	attemptLog, err := newAuthAttemptLog(dr, p)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := attemptLog.Insert(ctx, db, boil.Infer()); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func newAuthAttemptLog(dr DecodedAuthRequest, p auth.ClientPolicy) (*models.AuthAttemptLog, error) {
	serialized, err := json.Marshal(p.Custom)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.AuthAttemptLog{
		AuthAttemptLogClientType:            dr.ClientType,
		AuthAttemptLogClientInterface:       dr.ClientInterface,
		AuthAttemptLogClientIP:              dr.ClientIP,
		AuthAttemptLogClientMacAddress:      dr.ClientMAC,
		AuthAttemptLogClientGatewayName:     dr.GatewayName,
		AuthAttemptLogClientURL:             dr.GatewayURL,
		AuthAttemptLogClientHashID:          dr.HID,
		AuthAttemptLogOriginURL:             dr.OriginURL,
		AuthAttemptLogThemeSpecPath:         dr.ThemeSpec,
		AuthAttemptLogOpenndsVersion:        dr.Version,
		AuthAttemptLogRhid:                  p.ClientRHID,
		AuthAttemptLogSessionLengthMinutes:  int64(p.SessionDuration.Minutes()),
		AuthAttemptLogUploadRateThreshold:   p.UploadRateThresholdKBs,
		AuthAttemptLogDownloadRateThreshold: p.DownloadRateThresholdKBs,
		AuthAttemptLogUploadQuota:           p.UploadQuotaKB,
		AuthAttemptLogDownloadQuota:         p.DownloadQuotaKB,
		AuthAttemptLogCustomValue:           null.JSONFrom(serialized),
	}, nil
}

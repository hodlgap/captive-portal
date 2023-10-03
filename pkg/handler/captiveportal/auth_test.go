package captiveportal

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/go-redis/redismock/v9"
)

func mustAES256encode(plaintext, key, iv string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(errors.WithStack(err))
	}

	cipherText := make([]byte, len(plaintext))
	cbc := cipher.NewCBCEncrypter(block, []byte(iv))
	cbc.CryptBlocks(cipherText, []byte(plaintext))

	return base64.StdEncoding.EncodeToString(cipherText)
}

func toPlaintext(r DecodedAuthRequest) string {
	return strings.Join([]string{
		"authdir=" + r.AuthDir,
		"client_type=" + r.ClientType,
		"clientif=" + r.ClientInterface,
		"clientip=" + r.ClientIP,
		"clientmac=" + r.ClientMAC,
		"gatewayaddress=" + r.GatewayAddress,
		"gatewaymac=" + r.GatewayMAC,
		"gatewayname=" + r.GatewayName,
		"gatewayurl=" + r.GatewayURL,
		"hid=" + r.HID,
		"originurl=" + r.OriginURL,
		"themespec=" + r.ThemeSpec,
		"version=" + r.Version,
	}, ", ")
}

func TestAuthHandler_base_scenario(t *testing.T) {
	// This below struct is the plaintext of the FAS request
	_ = DecodedAuthRequest{
		AuthDir:         "opennds_auth",
		ClientType:      "cpi_url",
		ClientInterface: "br-lan",
		ClientIP:        "192.168.1.100",
		ClientMAC:       "f8:e4:3b:d6:1e:43",
		GatewayAddress:  "192.168.1.1:2050",
		GatewayMAC:      "88366c31a26c",
		GatewayName:     "gwgwgw%20Node%3a88366c31a26c%20",
		GatewayURL:      "http%3a%2f%2fstatus.client",
		HID:             "cd3b05fcb1bd99513cf1191f4972bcb053149cfe85798f6f5b199b348bae840b",
		OriginURL:       "http%3a%2f%2fstatus.client%2f",
		ThemeSpec:       "",
		Version:         "10.1.3",
	}
	const (
		// encryptedFAS is the FAS request encrypted with AES-256-CBC
		encryptedFAS  = "TFFQNzRFT2tBZTZDcEV4akhrZ0o2dDN3VkJMMEx1ZGllTWtQcG5ybGVUNHJkNW9acHI4UjFDRUtaZXdiTkoveXNMbFV4YUlEaXpIUGhTaWF1WjZyQndaY00wSHBKYjNmRkhRN0pmSEQwY2p6RDZYUlFnRlErMTRRTDNDQ3lFdFJ5ZE5Ud2NXekJYaWlVWUVtZDJiUmhieittdzZjRU1BZzJzREc1VTMyejZHN2lLR0VadE9qZmxTRjYyUWlBMnRoK1BnZVpyWS9IelkwZkZLTXBTN2RZTDJ3bDdva3hsK0pIVFV4S3dyWXFkeEFTTGUwQjRCTnBTVFpoZytkTkZITmdxVy9LdTg4Vzh1RXM0Qm1NeHUzS0pXT1c2T0VPcVNJV2pmWmlERkFLTFIxRUxlWmpBNE51VGdQQ3VDWU5jZ0lidEUvY2dZOFBHbFhVQW51eEZRRjcwMjBoQXluL3VHcUJTbFkwU0NHR2dqd3huYmtOS2ZxdVhVKzNEQzZ3Nlg4OEx5TVVzU0ZTd2d2WWZHRjJSbVZ2RVArSWVMVnlRaHphSDNFeUNVeDVaZGZ0bWowZ3k3aGtTVjBYaE1saEF5dzVvcUN4Rm1rSE5CWmRmNmNBcy9zY2JuZjhCMGVmTjRhdVZWd2NpV1FOMzB6T3B1V2F0Z3YvVzFOVGorbzdROWllMy9YVXFkR1NkclFQeE5uSldYWnZRU2V3TENFS216SzQwbWZUZVMzT1ZNcTU2TkxQS2gvRVBuWnhqeUx6ZVlr"
		iv            = "448b1bc37d4021f8"
		encryptionKey = "1234567890abcdef1234567890abcdef"

		expectedGatewayHash = "685ae1e926c631a63e381ae3de10ad8afa829cf0774eb0bbded06b8c1999f09d"
		expectedRHID        = "fc64942e8eef184be51b499d947adc2f65ab09f08b2c90adb07955290bfc9a93"

		// expectedClientSetting compiled as follows:
		// sessionlength := 1001 // 1001 minutes
		// uploadrate := 1002    // Upload Rate Limit Threshold: 1002 Kb/s
		// downloadrate := 1003  // Download Rate Limit Threshold: 1003 Kb/s
		// uploadquota := 1004   // Upload Quota: 1004 KBytes
		// downloadquota := 1005 // Download Quota: 1005 KBytes
		// custom := "key=value, key2=value2"
		expectedClientSetting = "fc64942e8eef184be51b499d947adc2f65ab09f08b2c90adb07955290bfc9a93%201001%201002%201003%201004%201005%20a2V5PXZhbHVlLCBrZXkyPXZhbHVlMg%3D%3D"
		expectedExpiration    = 0
	)

	req := httptest.NewRequest(http.MethodGet, AuthHandlerURL, nil)
	q := req.URL.Query()
	q.Set("fas", encryptedFAS)
	q.Set("iv", iv)
	req.URL.RawQuery = q.Encode()

	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	rCli, mock := redismock.NewClientMock()
	mock.ExpectLPush(expectedGatewayHash, expectedRHID).SetVal(1)
	mock.ExpectSet(
		fmt.Sprintf("%s/%s", expectedGatewayHash, expectedRHID),
		expectedClientSetting,
		expectedExpiration,
	).SetVal("OK")

	if assert.NoError(t, NewAuthHandler(encryptionKey, rCli)(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

package auth

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type ClientPolicy struct {
	ClientRHID               string
	SessionDuration          time.Duration // Session Duration: minutes
	UploadRateThresholdKBs   int64         // Upload Rate Limit Threshold: Kb/s
	DownloadRateThresholdKBs int64         // Download Rate Limit Threshold: Kb/s
	UploadQuotaKB            int64         // Upload Quota: KBytes
	DownloadQuotaKB          int64         // Download Quota: KBytes

	Custom map[string]any // Custom values
}

func (c *ClientPolicy) toAuthAttemptLog() {
	// TODO: implement
	panic("not implemented")
}

// toOpenNDSFormat converts ClientPolicy to the format that is expected by opennds auth daemon ("authmon")
func (c *ClientPolicy) toOpenNDSFormat() string {
	custom := ""
	if c.Custom != nil && len(c.Custom) > 0 {
		idx := 0
		kvFormats := make([]string, len(c.Custom))

		for k, v := range c.Custom {
			kvFormats[idx] = fmt.Sprintf("%s=%s", k, v)
			idx += 1
		}

		custom = strings.Join(kvFormats, ", ")
	}

	format := fmt.Sprintf(
		"%s %d %d %d %d %d %s",
		c.ClientRHID,
		int64(c.SessionDuration.Minutes()),
		c.UploadRateThresholdKBs,
		c.DownloadRateThresholdKBs,
		c.UploadQuotaKB,
		c.DownloadQuotaKB,
		base64.StdEncoding.EncodeToString([]byte(custom)),
	)

	format = url.QueryEscape(format)
	// NOTE(@Kcrong): url.QueryEscape replaces spaces with "+", but opennds needs "%20"
	format = strings.ReplaceAll(format, "+", "%20")

	return format
}

func (c *ClientPolicy) AddTag(key string, value any) {
	c.Custom[key] = value
}

func NewClientPolicy(
	rhid string,
	durationMinutes int64,
	uploadRate,
	downloadRate,
	uploadQuota,
	downloadQuota int64,
	tags map[string]any,
) *ClientPolicy {

	return &ClientPolicy{
		ClientRHID:               rhid,
		SessionDuration:          time.Duration(durationMinutes) * time.Minute,
		UploadRateThresholdKBs:   uploadRate,
		DownloadRateThresholdKBs: downloadRate,
		UploadQuotaKB:            uploadQuota,
		DownloadQuotaKB:          downloadQuota,
		Custom:                   tags,
	}
}

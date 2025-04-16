package etc

import (
	"fmt"
	"strings"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
)

var cfg = config.Load()

func AddImageBaseUrl(req string) string {
	if req != "" && !(strings.Contains(req, "http") || strings.Contains(req, "https")) {
		req = fmt.Sprintf("%s/%s", cfg.CdnImagesBucketBaseURL, req)
	}
	return req
}

func RemoveImageBaseUrl(req string) string {
	return strings.ReplaceAll(req, cfg.CdnImagesBucketBaseURL+"/", "")
}

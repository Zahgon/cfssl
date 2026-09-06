package sign

import (
	"net/http"

	"github.com/cloudflare/cfssl/config"
)

func NewHandler(caFile, caKeyFile string, policy *config.Signing) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func NewAuthHandler(caFile, caKeyFile string, policy *config.Signing) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

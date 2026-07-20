package universal

import (
	"github.com/cloudflare/cfssl/ocsp"
	ocspConfig "github.com/cloudflare/cfssl/ocsp/config"
)

func NewSignerFromConfig(cfg ocspConfig.Config) (ocsp.Signer, error) {
	_ = "STUB: not implemented"
	return *new(ocsp.Signer), nil
}

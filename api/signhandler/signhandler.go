package signhandler

import (
	"math/big"
	"net/http"

	"github.com/cloudflare/cfssl/api"
	"github.com/cloudflare/cfssl/bundler"
	"github.com/cloudflare/cfssl/signer"
)

const NoBundlerMessage = `This request requires a bundler, but one is not initialized for the API server.`

type Handler struct {
	signer  signer.Signer
	bundler *bundler.Bundler
}

func NewHandlerFromSigner(signer signer.Signer) (h *api.HTTPHandler, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Handler) SetBundler(caBundleFile, intBundleFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type jsonSignRequest struct {
	Hostname string          `json:"hostname"`
	Hosts    []string        `json:"hosts"`
	Request  string          `json:"certificate_request"`
	Subject  *signer.Subject `json:"subject,omitempty"`
	Profile  string          `json:"profile"`
	Label    string          `json:"label"`
	Serial   *big.Int        `json:"serial,omitempty"`
	Bundle   bool            `json:"bundle"`
}

func jsonReqToTrue(js jsonSignRequest) signer.SignRequest {
	_ = "STUB: not implemented"
	return *new(signer.SignRequest)
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

type AuthHandler struct {
	signer  signer.Signer
	bundler *bundler.Bundler
}

func NewAuthHandlerFromSigner(signer signer.Signer) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (h *AuthHandler) SetBundler(caBundleFile, intBundleFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (h *AuthHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

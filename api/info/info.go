package info

import (
	"net/http"

	"github.com/cloudflare/cfssl/signer"
)

type Handler struct {
	sign signer.Signer
}

func NewHandler(s signer.Signer) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

type MultiHandler struct {
	signers      map[string]signer.Signer
	defaultLabel string
}

func NewMultiHandler(signers map[string]signer.Signer, defaultLabel string) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (h *MultiHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

package crl

import (
	"crypto"
	"crypto/x509"
	"net/http"

	"github.com/cloudflare/cfssl/certdb"
)

type Handler struct {
	dbAccessor certdb.Accessor
	ca         *x509.Certificate
	key        crypto.Signer
}

func NewHandler(dbAccessor certdb.Accessor, caPath string, caKeyPath string) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

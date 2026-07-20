package ocsp

import (
	"crypto"
	"net/http"

	"github.com/cloudflare/cfssl/ocsp"
)

type Handler struct {
	signer ocsp.Signer
}

func NewHandler(s ocsp.Signer) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

type jsonSignRequest struct {
	Certificate string `json:"certificate"`
	Status      string `json:"status"`
	Reason      int    `json:"reason,omitempty"`
	RevokedAt   string `json:"revoked_at,omitempty"`
	IssuerHash  string `json:"issuer_hash,omitempty"`
}

var nameToHash = map[string]crypto.Hash{
	"MD5":    crypto.MD5,
	"SHA1":   crypto.SHA1,
	"SHA256": crypto.SHA256,
	"SHA384": crypto.SHA384,
	"SHA512": crypto.SHA512,
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

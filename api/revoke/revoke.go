package revoke

import (
	"net/http"

	"github.com/cloudflare/cfssl/certdb"
	"github.com/cloudflare/cfssl/ocsp"
)

type Handler struct {
	dbAccessor certdb.Accessor
	Signer     ocsp.Signer
}

func NewHandler(dbAccessor certdb.Accessor) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func NewOCSPHandler(dbAccessor certdb.Accessor, signer ocsp.Signer) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type jsonRevokeRequest struct {
	Serial string `json:"serial"`
	AKI    string `json:"authority_key_id"`
	Reason string `json:"reason"`
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

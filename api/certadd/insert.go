package certadd

import (
	"net/http"
	"time"

	"github.com/cloudflare/cfssl/certdb"
	"github.com/cloudflare/cfssl/ocsp"
	"github.com/jmoiron/sqlx/types"

	stdocsp "golang.org/x/crypto/ocsp"
)

type Handler struct {
	dbAccessor certdb.Accessor
	signer     ocsp.Signer
}

func NewHandler(dbAccessor certdb.Accessor, signer ocsp.Signer) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type AddRequest struct {
	Serial       string         `json:"serial_number"`
	AKI          string         `json:"authority_key_identifier"`
	CALabel      string         `json:"ca_label"`
	Status       string         `json:"status"`
	Reason       int            `json:"reason"`
	Expiry       time.Time      `json:"expiry"`
	RevokedAt    time.Time      `json:"revoked_at"`
	PEM          string         `json:"pem"`
	IssuedAt     *time.Time     `json:"issued_at"`
	NotBefore    *time.Time     `json:"not_before"`
	MetadataJSON types.JSONText `json:"metadata"`
	SansJSON     types.JSONText `json:"sans"`
	CommonName   string         `json:"common_name"`
}

var validReasons = map[int]bool{
	stdocsp.Unspecified:          true,
	stdocsp.KeyCompromise:        true,
	stdocsp.CACompromise:         true,
	stdocsp.AffiliationChanged:   true,
	stdocsp.Superseded:           true,
	stdocsp.CessationOfOperation: true,
	stdocsp.CertificateHold:      true,
	stdocsp.RemoveFromCRL:        true,
	stdocsp.PrivilegeWithdrawn:   true,
	stdocsp.AACompromise:         true,
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

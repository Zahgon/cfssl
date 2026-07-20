package certdb

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx/types"
)

type CertificateRecord struct {
	Serial    string    `db:"serial_number"`
	AKI       string    `db:"authority_key_identifier"`
	CALabel   string    `db:"ca_label"`
	Status    string    `db:"status"`
	Reason    int       `db:"reason"`
	Expiry    time.Time `db:"expiry"`
	RevokedAt time.Time `db:"revoked_at"`
	PEM       string    `db:"pem"`

	IssuedAt     *time.Time     `db:"issued_at"`
	NotBefore    *time.Time     `db:"not_before"`
	MetadataJSON types.JSONText `db:"metadata"`
	SANsJSON     types.JSONText `db:"sans"`
	CommonName   sql.NullString `db:"common_name"`
}

func (c *CertificateRecord) SetMetadata(meta map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CertificateRecord) GetMetadata() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CertificateRecord) SetSANs(meta []string) error { _ = "STUB: not implemented"; return nil }

func (c *CertificateRecord) GetSANs() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

type OCSPRecord struct {
	Serial string    `db:"serial_number"`
	AKI    string    `db:"authority_key_identifier"`
	Body   string    `db:"body"`
	Expiry time.Time `db:"expiry"`
}

type Accessor interface {
	InsertCertificate(cr CertificateRecord) error
	GetCertificate(serial, aki string) ([]CertificateRecord, error)
	GetUnexpiredCertificates() ([]CertificateRecord, error)
	GetRevokedAndUnexpiredCertificates() ([]CertificateRecord, error)
	GetUnexpiredCertificatesByLabel(labels []string) (crs []CertificateRecord, err error)
	GetRevokedAndUnexpiredCertificatesByLabel(label string) ([]CertificateRecord, error)
	GetRevokedAndUnexpiredCertificatesByLabelSelectColumns(label string) ([]CertificateRecord, error)
	RevokeCertificate(serial, aki string, reasonCode int) error
	InsertOCSP(rr OCSPRecord) error
	GetOCSP(serial, aki string) ([]OCSPRecord, error)
	GetUnexpiredOCSPs() ([]OCSPRecord, error)
	UpdateOCSP(serial, aki, body string, expiry time.Time) error
	UpsertOCSP(serial, aki, body string, expiry time.Time) error
}

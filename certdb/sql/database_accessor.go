package sql

import (
	"time"

	"github.com/cloudflare/cfssl/certdb"

	"github.com/jmoiron/sqlx"
	"github.com/kisielk/sqlstruct"
)

func init() {
	sqlstruct.TagName = "db"
}

const (
	insertSQL = `
INSERT INTO certificates (serial_number, authority_key_identifier, ca_label, status, reason, expiry, revoked_at, pem,
	issued_at, not_before, metadata, sans, common_name)
VALUES (:serial_number, :authority_key_identifier, :ca_label, :status, :reason, :expiry, :revoked_at, :pem,
	:issued_at, :not_before, :metadata, :sans, :common_name);`

	selectSQL = `
SELECT %s FROM certificates
	WHERE (serial_number = ? AND authority_key_identifier = ?);`

	selectAllUnexpiredSQL = `
SELECT %s FROM certificates
	WHERE CURRENT_TIMESTAMP < expiry;`

	selectAllRevokedAndUnexpiredWithLabelSQL = `
SELECT %s FROM certificates
	WHERE CURRENT_TIMESTAMP < expiry AND status='revoked' AND ca_label= ?;`

	selectRevokedAndUnexpiredWithLabelSQL = `
SELECT serial_number, revoked_at FROM certificates
	WHERE CURRENT_TIMESTAMP < expiry AND status='revoked' AND ca_label= ?;`

	selectAllRevokedAndUnexpiredSQL = `
SELECT %s FROM certificates
	WHERE CURRENT_TIMESTAMP < expiry AND status='revoked';`

	updateRevokeSQL = `
UPDATE certificates
	SET status='revoked', revoked_at=CURRENT_TIMESTAMP, reason=:reason
	WHERE (serial_number = :serial_number AND authority_key_identifier = :authority_key_identifier);`

	insertOCSPSQL = `
INSERT INTO ocsp_responses (serial_number, authority_key_identifier, body, expiry)
  VALUES (:serial_number, :authority_key_identifier, :body, :expiry);`

	updateOCSPSQL = `
UPDATE ocsp_responses
  SET body = :body, expiry = :expiry
	WHERE (serial_number = :serial_number AND authority_key_identifier = :authority_key_identifier);`

	selectAllUnexpiredOCSPSQL = `
SELECT %s FROM ocsp_responses
	WHERE CURRENT_TIMESTAMP < expiry;`

	selectOCSPSQL = `
SELECT %s FROM ocsp_responses
  WHERE (serial_number = ? AND authority_key_identifier = ?);`
)

type Accessor struct {
	db *sqlx.DB
}

var _ certdb.Accessor = &Accessor{}

func wrapSQLError(err error) error { _ = "STUB: not implemented"; return nil }

func (d *Accessor) checkDB() error { _ = "STUB: not implemented"; return nil }

func NewAccessor(db *sqlx.DB) *Accessor { _ = "STUB: not implemented"; return nil }

func (d *Accessor) SetDB(db *sqlx.DB) { _ = "STUB: not implemented"; return }

func (d *Accessor) InsertCertificate(cr certdb.CertificateRecord) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Accessor) GetCertificate(serial, aki string) (crs []certdb.CertificateRecord, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Accessor) GetUnexpiredCertificates() (crs []certdb.CertificateRecord, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Accessor) GetUnexpiredCertificatesByLabel(labels []string) (crs []certdb.CertificateRecord, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Accessor) GetRevokedAndUnexpiredCertificates() (crs []certdb.CertificateRecord, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Accessor) GetRevokedAndUnexpiredCertificatesByLabel(label string) (crs []certdb.CertificateRecord, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Accessor) GetRevokedAndUnexpiredCertificatesByLabelSelectColumns(label string) (crs []certdb.CertificateRecord, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Accessor) RevokeCertificate(serial, aki string, reasonCode int) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Accessor) InsertOCSP(rr certdb.OCSPRecord) error { _ = "STUB: not implemented"; return nil }

func (d *Accessor) GetOCSP(serial, aki string) (ors []certdb.OCSPRecord, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Accessor) GetUnexpiredOCSPs() (ors []certdb.OCSPRecord, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Accessor) UpdateOCSP(serial, aki, body string, expiry time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Accessor) UpsertOCSP(serial, aki, body string, expiry time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

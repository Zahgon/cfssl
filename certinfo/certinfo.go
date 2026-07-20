package certinfo

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"time"

	"github.com/cloudflare/cfssl/certdb"
)

type Certificate struct {
	Subject            Name      `json:"subject,omitempty"`
	Issuer             Name      `json:"issuer,omitempty"`
	SerialNumber       string    `json:"serial_number,omitempty"`
	SANs               []string  `json:"sans,omitempty"`
	NotBefore          time.Time `json:"not_before"`
	NotAfter           time.Time `json:"not_after"`
	SignatureAlgorithm string    `json:"sigalg"`
	AKI                string    `json:"authority_key_id"`
	SKI                string    `json:"subject_key_id"`
	RawPEM             string    `json:"pem"`
}

type Name struct {
	CommonName         string        `json:"common_name,omitempty"`
	SerialNumber       string        `json:"serial_number,omitempty"`
	Country            string        `json:"country,omitempty"`
	Organization       string        `json:"organization,omitempty"`
	OrganizationalUnit string        `json:"organizational_unit,omitempty"`
	Locality           string        `json:"locality,omitempty"`
	Province           string        `json:"province,omitempty"`
	StreetAddress      string        `json:"street_address,omitempty"`
	PostalCode         string        `json:"postal_code,omitempty"`
	Names              []interface{} `json:"names,omitempty"`
}

func ParseName(name pkix.Name) Name { _ = "STUB: not implemented"; return *new(Name) }

func formatKeyID(id []byte) string { _ = "STUB: not implemented"; return "" }

func ParseCertificate(cert *x509.Certificate) *Certificate { _ = "STUB: not implemented"; return nil }

func ParseCertificateFile(certFile string) (*Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCertificatePEM(certPEM []byte) (*Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCSRPEM(csrPEM []byte) (*x509.CertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCSRFile(csrFile string) (*x509.CertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCertificateDomain(domain string) (cert *Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseSerialNumber(serial, aki string, dbAccessor certdb.Accessor) (*Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

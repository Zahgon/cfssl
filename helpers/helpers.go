package helpers

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"time"

	ct "github.com/google/certificate-transparency-go"
	"golang.org/x/crypto/ocsp"
)

const OneYear = 8760 * time.Hour

const OneDay = 24 * time.Hour

var DelegationUsage = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 44363, 44}

var DelegationExtension = pkix.Extension{
	Id:       DelegationUsage,
	Critical: false,
	Value:    []byte{0x05, 0x00},
}

func InclusiveDate(year int, month time.Month, day int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

var Jul2012 = InclusiveDate(2012, time.July, 01)

var Apr2015 = InclusiveDate(2015, time.April, 01)

func KeyLength(key interface{}) int { _ = "STUB: not implemented"; return 0 }

func ExpiryTime(chain []*x509.Certificate) (notAfter time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func MonthsValid(c *x509.Certificate) int { _ = "STUB: not implemented"; return 0 }

func ValidExpiry(c *x509.Certificate) bool { _ = "STUB: not implemented"; return false }

func SignatureString(alg x509.SignatureAlgorithm) string { _ = "STUB: not implemented"; return "" }

func HashAlgoString(alg x509.SignatureAlgorithm) string { _ = "STUB: not implemented"; return "" }

func StringTLSVersion(version string) uint16 { _ = "STUB: not implemented"; return 0 }

func EncodeCertificatesPEM(certs []*x509.Certificate) []byte { _ = "STUB: not implemented"; return nil }

func EncodeCertificatePEM(cert *x509.Certificate) []byte { _ = "STUB: not implemented"; return nil }

func ParseCertificatesPEM(certsPEM []byte) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCertificatesDER(certsDER []byte, password string) (certs []*x509.Certificate, key crypto.Signer, err error) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.Signer), nil
}

func ParseSelfSignedCertificatePEM(certPEM []byte) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCertificatePEM(certPEM []byte) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseOneCertificateFromPEM(certsPEM []byte) ([]*x509.Certificate, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func LoadPEMCertPool(certsFile string) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PEMToCertPool(pemCerts []byte) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParsePrivateKeyPEM(keyPEM []byte) (key crypto.Signer, err error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func ParsePrivateKeyPEMWithPassword(keyPEM []byte, password []byte) (key crypto.Signer, err error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func GetKeyDERFromPEM(in []byte, password []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCSR(in []byte) (csr *x509.CertificateRequest, rest []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ParseCSRPEM(csrPEM []byte) (*x509.CertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCSRDER(csrDER []byte) (*x509.CertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SignerAlgo(priv crypto.Signer) x509.SignatureAlgorithm {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm)
}

func LoadClientCertificate(certFile string, keyFile string) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateTLSConfig(remoteCAs *x509.CertPool, cert *tls.Certificate) *tls.Config {
	_ = "STUB: not implemented"
	return nil
}

func SerializeSCTList(sctList []ct.SignedCertificateTimestamp) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeserializeSCTList(serializedSCTList []byte) ([]ct.SignedCertificateTimestamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SCTListFromOCSPResponse(response *ocsp.Response) ([]ct.SignedCertificateTimestamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadBytes(valFile string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

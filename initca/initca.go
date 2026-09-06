package initca

import (
	"crypto"
	"crypto/x509"

	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/helpers"
)

func validator(req *csr.CertificateRequest) error { _ = "STUB: not implemented"; return nil }

func New(req *csr.CertificateRequest) (cert, csrPEM, key []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func NewFromPEM(req *csr.CertificateRequest, keyFile string) (cert, csrPEM []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func RenewFromPEM(caFile, keyFile string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFromSigner(req *csr.CertificateRequest, priv crypto.Signer) (cert, csrPEM []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func RenewFromSigner(ca *x509.Certificate, priv crypto.Signer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var CAPolicy = func() *config.Signing {
	return &config.Signing{
		Default: &config.SigningProfile{
			Usage:        []string{"cert sign", "crl sign"},
			ExpiryString: "43800h",
			Expiry:       5 * helpers.OneYear,
			CAConstraint: config.CAConstraint{IsCA: true},
		},
	}
}

func Update(ca *x509.Certificate, priv crypto.Signer) (cert []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

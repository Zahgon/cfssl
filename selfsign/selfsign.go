package selfsign

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"time"

	"github.com/cloudflare/cfssl/config"
)

const threeMonths = 2190 * time.Hour

func parseCertificateRequest(priv crypto.Signer, csrBytes []byte) (template *x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type subjectPublicKeyInfo struct {
	Algorithm        pkix.AlgorithmIdentifier
	SubjectPublicKey asn1.BitString
}

func Sign(priv crypto.Signer, csrPEM []byte, profile *config.SigningProfile) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

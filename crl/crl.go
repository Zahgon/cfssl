package crl

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"time"

	"github.com/cloudflare/cfssl/certdb"
)

func NewCRLFromFile(serialList, issuerFile, keyFile []byte, expiryTime string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCRLFromDB(certs []certdb.CertificateRecord, issuerCert *x509.Certificate, key crypto.Signer, expiryTime time.Duration) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateGenericCRL(certList []pkix.RevokedCertificate, key crypto.Signer, issuingCert *x509.Certificate, expiryTime time.Time) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

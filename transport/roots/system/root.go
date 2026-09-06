package system

import (
	"crypto/x509"
)

func appendPEM(roots []*x509.Certificate, pemCerts []byte) ([]*x509.Certificate, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func New(metadata map[string]string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

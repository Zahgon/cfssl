package roots

import (
	"crypto/x509"

	"github.com/cloudflare/cfssl/transport/core"
	"github.com/cloudflare/cfssl/transport/roots/system"
)

var Providers = map[string]func(map[string]string) ([]*x509.Certificate, error){
	"system": system.New,
	"cfssl":  NewCFSSL,
	"file":   TrustPEM,
}

type TrustStore struct {
	roots map[string]*x509.Certificate
}

func (ts *TrustStore) Pool() *x509.CertPool { _ = "STUB: not implemented"; return nil }

func (ts *TrustStore) Certificates() []*x509.Certificate { _ = "STUB: not implemented"; return nil }

func (ts *TrustStore) addCerts(certs []*x509.Certificate) { _ = "STUB: not implemented"; return }

type Trusted interface {
	Certificates() []*x509.Certificate

	AddCert(cert *x509.Certificate)

	AddPEM(cert []byte) bool
}

func New(rootDefs []*core.Root) (*TrustStore, error) { _ = "STUB: not implemented"; return nil, nil }

func TrustPEM(metadata map[string]string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

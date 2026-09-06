package transport

import (
	"crypto/tls"
	"time"

	"github.com/cloudflare/backoff"
	"github.com/cloudflare/cfssl/transport/ca"
	"github.com/cloudflare/cfssl/transport/core"
	"github.com/cloudflare/cfssl/transport/kp"
	"github.com/cloudflare/cfssl/transport/roots"
)

func envOrDefault(key, def string) string { _ = "STUB: not implemented"; return "" }

var (
	NewKeyProvider = func(id *core.Identity) (kp.KeyProvider, error) {
		return kp.NewStandardProvider(id)
	}

	NewCA = func(id *core.Identity) (ca.CertificateAuthority, error) {
		return ca.NewCFSSLProvider(id, nil)
	}
)

type Transport struct {
	Before time.Duration

	Provider kp.KeyProvider

	CA ca.CertificateAuthority

	TrustStore *roots.TrustStore

	ClientTrustStore *roots.TrustStore

	Identity *core.Identity

	Backoff *backoff.Backoff

	RevokeSoftFail bool
}

func (tr *Transport) TLSClientAuthClientConfig(host string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tr *Transport) TLSClientAuthServerConfig() (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tr *Transport) TLSServerConfig() (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(before time.Duration, identity *core.Identity) (*Transport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tr *Transport) Lifespan() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (tr *Transport) RefreshKeys() (err error) { _ = "STUB: not implemented"; return nil }

func (tr *Transport) getCertificate() (cert tls.Certificate, err error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

func Dial(address string, tr *Transport) (*tls.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tr *Transport) AutoUpdate(certUpdates chan<- time.Time, errChan chan<- error) {
	_ = "STUB: not implemented"
	return
}

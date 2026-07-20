package kp

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"errors"

	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/transport/core"
)

const (
	curveP256 = 256
	curveP384 = 384
	curveP521 = 521
)

type KeyProvider interface {
	Certificate() *x509.Certificate

	CertificateRequest(*csr.CertificateRequest) ([]byte, error)

	Check() error

	Generate(algo string, size int) error

	Load() error

	Persistent() bool

	Ready() bool

	SetCertificatePEM([]byte) error

	SignalFailure(err error) bool

	SignCSR(csr *x509.CertificateRequest) ([]byte, error)

	Store() error

	X509KeyPair() (tls.Certificate, error)
}

type StandardPaths struct {
	KeyFile  string `json:"private_key"`
	CertFile string `json:"certificate"`
}

type StandardProvider struct {
	Paths    StandardPaths `json:"paths"`
	internal struct {
		priv crypto.Signer
		cert *x509.Certificate

		keyPEM  []byte
		certPEM []byte
	}
}

func NewStandardProvider(id *core.Identity) (*StandardProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sp *StandardProvider) resetCert() { _ = "STUB: not implemented"; return }

func (sp *StandardProvider) resetKey() { _ = "STUB: not implemented"; return }

var (
	ErrMissingKeyPath = errors.New("transport: standard provider is missing a private key path to accompany the certificate path")

	ErrMissingCertPath = errors.New("transport: standard provider is missing a certificate path to accompany the certificate path")
)

func (sp *StandardProvider) Check() error { _ = "STUB: not implemented"; return nil }

func (sp *StandardProvider) Persistent() bool { _ = "STUB: not implemented"; return false }

func (sp *StandardProvider) Generate(algo string, size int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (sp *StandardProvider) Certificate() *x509.Certificate { _ = "STUB: not implemented"; return nil }

func (sp *StandardProvider) CertificateRequest(req *csr.CertificateRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var ErrCertificateUnavailable = errors.New("transport: certificate unavailable")

func (sp *StandardProvider) Load() (err error) { _ = "STUB: not implemented"; return nil }

func (sp *StandardProvider) Ready() bool { _ = "STUB: not implemented"; return false }

func (sp *StandardProvider) SetCertificatePEM(certPEM []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (sp *StandardProvider) SignalFailure(err error) bool { _ = "STUB: not implemented"; return false }

func (sp *StandardProvider) SignCSR(tpl *x509.CertificateRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sp *StandardProvider) Store() error { _ = "STUB: not implemented"; return nil }

func (sp *StandardProvider) X509KeyPair() (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

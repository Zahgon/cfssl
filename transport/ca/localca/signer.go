package localca

import (
	"errors"

	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/signer/local"
)

type CA struct {
	s        *local.Signer
	disabled bool

	Label   string `json:"label"`
	Profile string `json:"profile"`

	KeyFile  string `json:"private_key,omitempty"`
	CertFile string `json:"certificate,omitempty"`
}

func (lca *CA) Toggle() { _ = "STUB: not implemented"; return }

var errNotSetup = errors.New("transport: local CA has not been setup")

func (lca *CA) CACertificate() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

var errDisabled = errors.New("transport: local CA is deactivated")

func (lca *CA) SignCSR(csrPEM []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ExampleRequest() *csr.CertificateRequest { _ = "STUB: not implemented"; return nil }

func ExampleSigningConfig() *config.Signing { _ = "STUB: not implemented"; return nil }

func New(req *csr.CertificateRequest, profiles *config.Signing) (*CA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFromSigner(s *local.Signer) *CA { _ = "STUB: not implemented"; return nil }

func Load(lca *CA, profiles *config.Signing) (err error) { _ = "STUB: not implemented"; return nil }

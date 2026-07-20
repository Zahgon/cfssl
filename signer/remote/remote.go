package remote

import (
	"crypto/x509"
	"net/http"

	"github.com/cloudflare/cfssl/certdb"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/info"
	"github.com/cloudflare/cfssl/signer"
)

type Signer struct {
	policy      *config.Signing
	reqModifier func(*http.Request, []byte)
}

func NewSigner(policy *config.Signing) (*Signer, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Signer) Sign(req signer.SignRequest) (cert []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) Info(req info.Req) (resp *info.Resp, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) remoteOp(req interface{}, profile, target string) (resp interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) SigAlgo() x509.SignatureAlgorithm {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm)
}

func (s *Signer) SetPolicy(policy *config.Signing) { _ = "STUB: not implemented"; return }

func (s *Signer) SetDBAccessor(dba certdb.Accessor) { _ = "STUB: not implemented"; return }

func (s *Signer) GetDBAccessor() certdb.Accessor {
	_ = "STUB: not implemented"
	return *new(certdb.Accessor)
}

func (s *Signer) SetReqModifier(mod func(*http.Request, []byte)) { _ = "STUB: not implemented"; return }

func (s *Signer) Policy() *config.Signing { _ = "STUB: not implemented"; return nil }

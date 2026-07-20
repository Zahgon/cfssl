package universal

import (
	"crypto/x509"
	"net/http"

	"github.com/cloudflare/cfssl/certdb"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/info"
	"github.com/cloudflare/cfssl/signer"
)

type Signer struct {
	local  signer.Signer
	remote signer.Signer
	policy *config.Signing
}

type Root struct {
	Config      map[string]string
	ForceRemote bool
}

type localSignerCheck func(root *Root, policy *config.Signing) (signer.Signer, bool, error)

func fileBackedSigner(root *Root, policy *config.Signing) (signer.Signer, bool, error) {
	_ = "STUB: not implemented"
	return *new(signer.Signer), false, nil
}

var localSignerList = []localSignerCheck{
	fileBackedSigner,
}

func PrependLocalSignerToList(signer localSignerCheck) { _ = "STUB: not implemented"; return }

func newLocalSigner(root Root, policy *config.Signing) (s signer.Signer, err error) {
	_ = "STUB: not implemented"
	return *new(signer.Signer), nil
}

func newUniversalSigner(root Root, policy *config.Signing) (*Signer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSigner(root Root, policy *config.Signing) (signer.Signer, error) {
	_ = "STUB: not implemented"
	return *new(signer.Signer), nil
}

func (s *Signer) getMatchingProfile(profile string) (*config.SigningProfile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) Sign(req signer.SignRequest) (cert []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) Info(req info.Req) (resp *info.Resp, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) SetDBAccessor(dba certdb.Accessor) { _ = "STUB: not implemented"; return }

func (s *Signer) GetDBAccessor() certdb.Accessor {
	_ = "STUB: not implemented"
	return *new(certdb.Accessor)
}

func (s *Signer) SetReqModifier(mod func(*http.Request, []byte)) { _ = "STUB: not implemented"; return }

func (s *Signer) SigAlgo() x509.SignatureAlgorithm {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm)
}

func (s *Signer) SetPolicy(policy *config.Signing) { _ = "STUB: not implemented"; return }

func (s *Signer) Policy() *config.Signing { _ = "STUB: not implemented"; return nil }

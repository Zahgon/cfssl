package local

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"net/http"

	"github.com/cloudflare/cfssl/certdb"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/info"
	"github.com/cloudflare/cfssl/signer"
	ct "github.com/google/certificate-transparency-go"

	"github.com/zmap/zlint/v3/lint"
)

type Signer struct {
	ca   *x509.Certificate
	priv crypto.Signer

	lintPriv   crypto.Signer
	policy     *config.Signing
	sigAlgo    x509.SignatureAlgorithm
	dbAccessor certdb.Accessor
}

func NewSigner(priv crypto.Signer, cert *x509.Certificate, sigAlgo x509.SignatureAlgorithm, policy *config.Signing) (*Signer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSignerFromFile(caFile, caKeyFile string, policy *config.Signing) (*Signer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type LintError struct {
	ErrorResults map[string]lint.LintResult
}

func (e *LintError) Error() string { _ = "STUB: not implemented"; return "" }

func (s *Signer) lint(template x509.Certificate, errLevel lint.LintStatus, lintRegistry lint.Registry) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Signer) sign(template *x509.Certificate, lintErrLevel lint.LintStatus, lintRegistry lint.Registry) (cert []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func replaceSliceIfEmpty(replaced, newContents *[]string) { _ = "STUB: not implemented"; return }

func PopulateSubjectFromCSR(s *signer.Subject, req pkix.Name) pkix.Name {
	_ = "STUB: not implemented"
	return *new(pkix.Name)
}

func OverrideHosts(template *x509.Certificate, hosts []string) { _ = "STUB: not implemented"; return }

func (s *Signer) Sign(req signer.SignRequest) (cert []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) SignFromPrecert(precert *x509.Certificate, scts []ct.SignedCertificateTimestamp) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) Info(req info.Req) (resp *info.Resp, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) SigAlgo() x509.SignatureAlgorithm {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm)
}

func (s *Signer) Certificate(label, profile string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) SetPolicy(policy *config.Signing) { _ = "STUB: not implemented"; return }

func (s *Signer) SetDBAccessor(dba certdb.Accessor) { _ = "STUB: not implemented"; return }

func (s *Signer) GetDBAccessor() certdb.Accessor {
	_ = "STUB: not implemented"
	return *new(certdb.Accessor)
}

func (s *Signer) SetReqModifier(func(*http.Request, []byte)) { _ = "STUB: not implemented"; return }

func (s *Signer) Policy() *config.Signing { _ = "STUB: not implemented"; return nil }

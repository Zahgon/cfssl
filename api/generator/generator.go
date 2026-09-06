package generator

import (
	"net/http"

	"github.com/cloudflare/cfssl/bundler"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/signer"
)

const (
	CSRNoHostMessage = `This certificate lacks a "hosts" field. This makes it unsuitable for
websites. For more information see the Baseline Requirements for the Issuance and Management
of Publicly-Trusted Certificates, v.1.1.6, from the CA/Browser Forum (https://cabforum.org);
specifically, section 10.2.3 ("Information Requirements").`

	NoBundlerMessage = `This request requires a bundler, but one is not initialized for the API server.`
)

type Sum struct {
	MD5    string `json:"md5"`
	SHA1   string `json:"sha-1"`
	SHA256 string `json:"sha-256"`
}

type Validator func(*csr.CertificateRequest) error

type CertRequest struct {
	Key  string         `json:"private_key"`
	CSR  string         `json:"certificate_request"`
	Sums map[string]Sum `json:"sums"`
}

type Handler struct {
	generator *csr.Generator
}

func NewHandler(validator Validator) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func computeSum(in []byte) (sum Sum, err error) { _ = "STUB: not implemented"; return *new(Sum), nil }

func (g *Handler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

type CertGeneratorHandler struct {
	generator *csr.Generator
	bundler   *bundler.Bundler
	signer    signer.Signer
}

func NewCertGeneratorHandler(validator Validator, caFile, caKeyFile string, policy *config.Signing) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func NewCertGeneratorHandlerFromSigner(validator Validator, signer signer.Signer) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (cg *CertGeneratorHandler) SetBundler(caBundleFile, intBundleFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type genSignRequest struct {
	Request *csr.CertificateRequest `json:"request"`
	Profile string                  `json:"profile"`
	Label   string                  `json:"label"`
	Bundle  bool                    `json:"bundle"`
}

func (cg *CertGeneratorHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func CSRValidate(req *csr.CertificateRequest) error { _ = "STUB: not implemented"; return nil }

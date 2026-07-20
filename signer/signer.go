package signer

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"math/big"
	"net/http"
	"time"

	"github.com/cloudflare/cfssl/certdb"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/info"
)

type Subject struct {
	CN           string
	Names        []csr.Name `json:"names"`
	SerialNumber string
}

type Extension struct {
	ID       config.OID `json:"id"`
	Critical bool       `json:"critical"`
	Value    string     `json:"value"`
}

type SignRequest struct {
	Hosts       []string    `json:"hosts"`
	Request     string      `json:"certificate_request"`
	Subject     *Subject    `json:"subject,omitempty"`
	Profile     string      `json:"profile"`
	CRLOverride string      `json:"crl_override"`
	Label       string      `json:"label"`
	Serial      *big.Int    `json:"serial,omitempty"`
	Extensions  []Extension `json:"extensions,omitempty"`

	NotBefore time.Time

	NotAfter time.Time

	ReturnPrecert bool

	Metadata map[string]interface{} `json:"metadata"`
}

func appendIf(s string, a *[]string) { _ = "STUB: not implemented"; return }

func (s *Subject) Name() pkix.Name { _ = "STUB: not implemented"; return *new(pkix.Name) }

func SplitHosts(hostList string) []string { _ = "STUB: not implemented"; return nil }

type Signer interface {
	Info(info.Req) (*info.Resp, error)
	Policy() *config.Signing
	SetDBAccessor(certdb.Accessor)
	GetDBAccessor() certdb.Accessor
	SetPolicy(*config.Signing)
	SigAlgo() x509.SignatureAlgorithm
	Sign(req SignRequest) (cert []byte, err error)
	SetReqModifier(func(*http.Request, []byte))
}

func Profile(s Signer, profile string) (*config.SigningProfile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DefaultSigAlgo(priv crypto.Signer) x509.SignatureAlgorithm {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm)
}

func isCommonAttr(t []int) bool { _ = "STUB: not implemented"; return false }

var caManagedExtensionOIDs = map[string]bool{

	asn1.ObjectIdentifier{2, 5, 29, 15}.String(): true,

	asn1.ObjectIdentifier{2, 5, 29, 37}.String(): true,

	asn1.ObjectIdentifier{2, 5, 29, 19}.String(): true,

	asn1.ObjectIdentifier{2, 5, 29, 14}.String(): true,

	asn1.ObjectIdentifier{2, 5, 29, 35}.String(): true,

	asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 1}.String(): true,

	asn1.ObjectIdentifier{2, 5, 29, 31}.String(): true,

	asn1.ObjectIdentifier{2, 5, 29, 32}.String(): true,

	asn1.ObjectIdentifier{2, 5, 29, 30}.String(): true,

	asn1.ObjectIdentifier{2, 5, 29, 17}.String(): true,

	asn1.ObjectIdentifier{2, 5, 29, 18}.String(): true,
}

func isCaManagedExtension(oid asn1.ObjectIdentifier) bool { _ = "STUB: not implemented"; return false }

func ParseCertificateRequest(s Signer, p *config.SigningProfile, csrBytes []byte) (template *x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type subjectPublicKeyInfo struct {
	Algorithm        pkix.AlgorithmIdentifier
	SubjectPublicKey asn1.BitString
}

func ComputeSKI(template *x509.Certificate) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FillTemplate(template *x509.Certificate, defaultProfile, profile *config.SigningProfile, notBefore time.Time, notAfter time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type policyInformation struct {
	PolicyIdentifier asn1.ObjectIdentifier
	Qualifiers       []interface{} `asn1:"tag:optional,omitempty"`
}

type cpsPolicyQualifier struct {
	PolicyQualifierID asn1.ObjectIdentifier
	Qualifier         string `asn1:"tag:optional,ia5"`
}

type userNotice struct {
	ExplicitText string `asn1:"tag:optional,utf8"`
}
type userNoticePolicyQualifier struct {
	PolicyQualifierID asn1.ObjectIdentifier
	Qualifier         userNotice
}

var (
	iDQTCertificationPracticeStatement = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 2, 1}

	iDQTUserNotice = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 2, 2}

	CTPoisonOID = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 4, 3}

	SCTListOID = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 4, 2}
)

func addPolicies(template *x509.Certificate, policies []config.CertificatePolicy) error {
	_ = "STUB: not implemented"
	return nil
}

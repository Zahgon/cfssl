package csr

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
)

const (
	curveP256 = 256
	curveP384 = 384
	curveP521 = 521
)

type Name struct {
	C            string            `json:"C,omitempty" yaml:"C,omitempty"`
	ST           string            `json:"ST,omitempty" yaml:"ST,omitempty"`
	L            string            `json:"L,omitempty" yaml:"L,omitempty"`
	O            string            `json:"O,omitempty" yaml:"O,omitempty"`
	OU           string            `json:"OU,omitempty" yaml:"OU,omitempty"`
	E            string            `json:"E,omitempty" yaml:"E,omitempty"`
	SerialNumber string            `json:"SerialNumber,omitempty" yaml:"SerialNumber,omitempty"`
	OID          map[string]string `json:"OID,omitempty", yaml:"OID,omitempty"`
}

type KeyRequest struct {
	A string `json:"algo" yaml:"algo"`
	S int    `json:"size" yaml:"size"`
}

func NewKeyRequest() *KeyRequest { _ = "STUB: not implemented"; return nil }

func (kr *KeyRequest) Algo() string { _ = "STUB: not implemented"; return "" }

func (kr *KeyRequest) Size() int { _ = "STUB: not implemented"; return 0 }

func (kr *KeyRequest) Generate() (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

func (kr *KeyRequest) SigAlgo() x509.SignatureAlgorithm {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm)
}

type CAConfig struct {
	PathLength  int    `json:"pathlen" yaml:"pathlen"`
	PathLenZero bool   `json:"pathlenzero" yaml:"pathlenzero"`
	Expiry      string `json:"expiry" yaml:"expiry"`
	Backdate    string `json:"backdate" yaml:"backdate"`
}

type CertificateRequest struct {
	CN                string           `json:"CN" yaml:"CN"`
	Names             []Name           `json:"names" yaml:"names"`
	Hosts             []string         `json:"hosts" yaml:"hosts"`
	KeyRequest        *KeyRequest      `json:"key,omitempty" yaml:"key,omitempty"`
	CA                *CAConfig        `json:"ca,omitempty" yaml:"ca,omitempty"`
	SerialNumber      string           `json:"serialnumber,omitempty" yaml:"serialnumber,omitempty"`
	DelegationEnabled bool             `json:"delegation_enabled,omitempty" yaml:"delegation_enabled,omitempty"`
	Extensions        []pkix.Extension `json:"extensions,omitempty" yaml:"extensions,omitempty"`
	CRL               string           `json:"crl_url,omitempty" yaml:"crl_url,omitempty"`
}

func New() *CertificateRequest { _ = "STUB: not implemented"; return nil }

func appendIf(s string, a *[]string) { _ = "STUB: not implemented"; return }

func OIDFromString(s string) (asn1.ObjectIdentifier, error) {
	_ = "STUB: not implemented"
	return *new(asn1.ObjectIdentifier), nil
}

func (cr *CertificateRequest) Name() (pkix.Name, error) {
	_ = "STUB: not implemented"
	return *new(pkix.Name), nil
}

type BasicConstraints struct {
	IsCA       bool `asn1:"optional"`
	MaxPathLen int  `asn1:"optional,default:-1"`
}

func ParseRequest(req *CertificateRequest) (csr, key []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ExtractCertificateRequest(cert *x509.Certificate) *CertificateRequest {
	_ = "STUB: not implemented"
	return nil
}

func getHosts(cert *x509.Certificate) []string { _ = "STUB: not implemented"; return nil }

func getNames(sub pkix.Name) []Name { _ = "STUB: not implemented"; return nil }

type Generator struct {
	Validator func(*CertificateRequest) error
}

func (g *Generator) ProcessRequest(req *CertificateRequest) (csr, key []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func IsNameEmpty(n Name) bool { _ = "STUB: not implemented"; return false }

func Regenerate(priv crypto.Signer, csr []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GenerateDER(priv crypto.Signer, req *CertificateRequest) (csr []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Generate(priv crypto.Signer, req *CertificateRequest) (csr []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendCAInfoToCSR(reqConf *CAConfig, csr *x509.CertificateRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func appendExtensionsToCSR(extensions []pkix.Extension, csr *x509.CertificateRequest) error {
	_ = "STUB: not implemented"
	return nil
}

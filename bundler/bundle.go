package bundler

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"time"
)

type Bundle struct {
	Chain       []*x509.Certificate
	Cert        *x509.Certificate
	Root        *x509.Certificate
	Key         interface{}
	Issuer      *pkix.Name
	Subject     *pkix.Name
	Expires     time.Time
	LeafExpires time.Time
	Hostnames   []string
	Status      *BundleStatus
}

type BundleStatus struct {
	IsRebundled bool `json:"rebundled"`

	ExpiringSKIs []string `json:"expiring_SKIs"`

	Untrusted []string `json:"untrusted_root_stores"`

	Messages []string `json:"messages"`

	Code int `json:"code"`
}

type chain []*x509.Certificate

func (c chain) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func PemBlockToString(block *pem.Block) string { _ = "STUB: not implemented"; return "" }

var typeToName = map[int]string{
	3:  "CommonName",
	5:  "SerialNumber",
	6:  "Country",
	7:  "Locality",
	8:  "Province",
	9:  "StreetAddress",
	10: "Organization",
	11: "OrganizationalUnit",
	17: "PostalCode",
}

type names []pkix.AttributeTypeAndValue

func (n names) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *Bundle) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *Bundle) buildHostnames() { _ = "STUB: not implemented"; return }

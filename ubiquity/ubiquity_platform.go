package ubiquity

import (
	"crypto/x509"
)

func SHA1RawPublicKey(cert *x509.Certificate) string { _ = "STUB: not implemented"; return "" }

type CertSet map[string]bool

func (s CertSet) Lookup(cert *x509.Certificate) bool { _ = "STUB: not implemented"; return false }

func (s CertSet) Add(cert *x509.Certificate) { _ = "STUB: not implemented"; return }

type Platform struct {
	Name            string `json:"name"`
	Weight          int    `json:"weight"`
	HashAlgo        string `json:"hash_algo"`
	KeyAlgo         string `json:"key_algo"`
	KeyStoreFile    string `json:"keystore"`
	KeyStore        CertSet
	HashUbiquity    HashUbiquity
	KeyAlgoUbiquity KeyAlgoUbiquity
}

func (p Platform) Trust(root *x509.Certificate) bool { _ = "STUB: not implemented"; return false }

func (p Platform) hashUbiquity() HashUbiquity { _ = "STUB: not implemented"; return *new(HashUbiquity) }

func (p Platform) keyAlgoUbiquity() KeyAlgoUbiquity {
	_ = "STUB: not implemented"
	return *new(KeyAlgoUbiquity)
}

func (p *Platform) ParseAndLoad() (ok bool) { _ = "STUB: not implemented"; return false }

var Platforms []Platform

func LoadPlatforms(filename string) error { _ = "STUB: not implemented"; return nil }

func UntrustedPlatforms(root *x509.Certificate) []string { _ = "STUB: not implemented"; return nil }

func CrossPlatformUbiquity(chain []*x509.Certificate) int { _ = "STUB: not implemented"; return 0 }

func ComparePlatformUbiquity(chain1, chain2 []*x509.Certificate) int {
	_ = "STUB: not implemented"
	return 0
}

func SHA2Homogeneity(chain []*x509.Certificate) int { _ = "STUB: not implemented"; return 0 }

func CompareSHA2Homogeneity(chain1, chain2 []*x509.Certificate) int {
	_ = "STUB: not implemented"
	return 0
}

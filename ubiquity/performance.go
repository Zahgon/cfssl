package ubiquity

import (
	"crypto/x509"
	"time"
)

func hashPriority(cert *x509.Certificate) int { _ = "STUB: not implemented"; return 0 }

func keyAlgoPriority(cert *x509.Certificate) int { _ = "STUB: not implemented"; return 0 }

func HashPriority(certs []*x509.Certificate) int { _ = "STUB: not implemented"; return 0 }

func KeyAlgoPriority(certs []*x509.Certificate) int { _ = "STUB: not implemented"; return 0 }

func CompareChainHashPriority(chain1, chain2 []*x509.Certificate) int {
	_ = "STUB: not implemented"
	return 0
}

func CompareChainKeyAlgoPriority(chain1, chain2 []*x509.Certificate) int {
	_ = "STUB: not implemented"
	return 0
}

func CompareChainCryptoSuite(chain1, chain2 []*x509.Certificate) int {
	_ = "STUB: not implemented"
	return 0
}

func CompareChainLength(chain1, chain2 []*x509.Certificate) int {
	_ = "STUB: not implemented"
	return 0
}

func compareTime(t1, t2 time.Time) int { _ = "STUB: not implemented"; return 0 }

func CompareChainExpiry(chain1, chain2 []*x509.Certificate) int {
	_ = "STUB: not implemented"
	return 0
}

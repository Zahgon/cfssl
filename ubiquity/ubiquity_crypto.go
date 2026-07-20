package ubiquity

import (
	"crypto/x509"
)

type HashUbiquity int

type KeyAlgoUbiquity int

const (
	UnknownHashUbiquity HashUbiquity = 0
	SHA2Ubiquity        HashUbiquity = 70
	SHA1Ubiquity        HashUbiquity = 100
	MD5Ubiquity         HashUbiquity = 0
	MD2Ubiquity         HashUbiquity = 0
)

const (
	RSAUbiquity         KeyAlgoUbiquity = 100
	DSAUbiquity         KeyAlgoUbiquity = 100
	ECDSA256Ubiquity    KeyAlgoUbiquity = 70
	ECDSA384Ubiquity    KeyAlgoUbiquity = 70
	ECDSA521Ubiquity    KeyAlgoUbiquity = 30
	UnknownAlgoUbiquity KeyAlgoUbiquity = 0
)

func hashUbiquity(cert *x509.Certificate) HashUbiquity {
	_ = "STUB: not implemented"
	return *new(HashUbiquity)
}

func keyAlgoUbiquity(cert *x509.Certificate) KeyAlgoUbiquity {
	_ = "STUB: not implemented"
	return *new(KeyAlgoUbiquity)
}

func ChainHashUbiquity(chain []*x509.Certificate) HashUbiquity {
	_ = "STUB: not implemented"
	return *new(HashUbiquity)
}

func ChainKeyAlgoUbiquity(chain []*x509.Certificate) KeyAlgoUbiquity {
	_ = "STUB: not implemented"
	return *new(KeyAlgoUbiquity)
}

func CompareChainHashUbiquity(chain1, chain2 []*x509.Certificate) int {
	_ = "STUB: not implemented"
	return 0
}

func CompareChainKeyAlgoUbiquity(chain1, chain2 []*x509.Certificate) int {
	_ = "STUB: not implemented"
	return 0
}

func CompareExpiryUbiquity(chain1, chain2 []*x509.Certificate) int {
	_ = "STUB: not implemented"
	return 0
}

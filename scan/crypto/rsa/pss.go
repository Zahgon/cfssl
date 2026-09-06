package rsa

import (
	"crypto"
	"hash"
	"io"
)

func emsaPSSEncode(mHash []byte, emBits int, salt []byte, hash hash.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func emsaPSSVerify(mHash, em []byte, emBits, sLen int, hash hash.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func signPSSWithSalt(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed, salt []byte) (s []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	PSSSaltLengthAuto = 0

	PSSSaltLengthEqualsHash = -1
)

type PSSOptions struct {
	SaltLength int

	Hash crypto.Hash
}

func (pssOpts *PSSOptions) HashFunc() crypto.Hash {
	_ = "STUB: not implemented"
	return *new(crypto.Hash)
}

func (opts *PSSOptions) saltLength() int { _ = "STUB: not implemented"; return 0 }

func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte, opts *PSSOptions) (s []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func VerifyPSS(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte, opts *PSSOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyPSS(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte, saltLen int) error {
	_ = "STUB: not implemented"
	return nil
}

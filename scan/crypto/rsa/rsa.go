package rsa

import (
	"crypto"
	"errors"
	"hash"
	"io"
	"math/big"
)

var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)

type PublicKey struct {
	N *big.Int
	E int
}

type OAEPOptions struct {
	Hash crypto.Hash

	Label []byte
}

var (
	errPublicModulus       = errors.New("crypto/rsa: missing public modulus")
	errPublicExponentSmall = errors.New("crypto/rsa: public exponent too small")
	errPublicExponentLarge = errors.New("crypto/rsa: public exponent too large")
)

func checkPub(pub *PublicKey) error { _ = "STUB: not implemented"; return nil }

type PrivateKey struct {
	PublicKey
	D      *big.Int
	Primes []*big.Int

	Precomputed PrecomputedValues
}

func (priv *PrivateKey) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

func (priv *PrivateKey) Sign(rand io.Reader, msg []byte, opts crypto.SignerOpts) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (priv *PrivateKey) Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PrecomputedValues struct {
	Dp, Dq *big.Int
	Qinv   *big.Int

	CRTValues []CRTValue
}

type CRTValue struct {
	Exp   *big.Int
	Coeff *big.Int
	R     *big.Int
}

func (priv *PrivateKey) Validate() error { _ = "STUB: not implemented"; return nil }

func GenerateKey(random io.Reader, bits int) (priv *PrivateKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (priv *PrivateKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func incCounter(c *[4]byte) { _ = "STUB: not implemented"; return }

func mgf1XOR(out []byte, hash hash.Hash, seed []byte) { _ = "STUB: not implemented"; return }

var ErrMessageTooLong = errors.New("crypto/rsa: message too long for RSA public key size")

func encrypt(c *big.Int, pub *PublicKey, m *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) (out []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var ErrDecryption = errors.New("crypto/rsa: decryption error")

var ErrVerification = errors.New("crypto/rsa: verification error")

func modInverse(a, n *big.Int) (ia *big.Int, ok bool) { _ = "STUB: not implemented"; return nil, false }

func (priv *PrivateKey) Precompute() { _ = "STUB: not implemented"; return }

func decrypt(random io.Reader, priv *PrivateKey, c *big.Int) (m *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptAndCheck(random io.Reader, priv *PrivateKey, c *big.Int) (m *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) (msg []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func leftPad(input []byte, size int) (out []byte) { _ = "STUB: not implemented"; return nil }

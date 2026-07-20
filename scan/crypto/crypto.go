package crypto

import (
	"hash"
	"io"
)

type Hash uint

func (h Hash) HashFunc() Hash { _ = "STUB: not implemented"; return *new(Hash) }

const (
	MD4 Hash = 1 + iota
	MD5
	SHA1
	SHA224
	SHA256
	SHA384
	SHA512
	MD5SHA1
	RIPEMD160
	SHA3_224
	SHA3_256
	SHA3_384
	SHA3_512
	SHA512_224
	SHA512_256
	maxHash
)

var digestSizes = []uint8{
	MD4:        16,
	MD5:        16,
	SHA1:       20,
	SHA224:     28,
	SHA256:     32,
	SHA384:     48,
	SHA512:     64,
	SHA512_224: 28,
	SHA512_256: 32,
	SHA3_224:   28,
	SHA3_256:   32,
	SHA3_384:   48,
	SHA3_512:   64,
	MD5SHA1:    36,
	RIPEMD160:  20,
}

func (h Hash) Size() int { _ = "STUB: not implemented"; return 0 }

var hashes = make([]func() hash.Hash, maxHash)

func (h Hash) New() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func (h Hash) Available() bool { _ = "STUB: not implemented"; return false }

func RegisterHash(h Hash, f func() hash.Hash) { _ = "STUB: not implemented"; return }

type PublicKey interface{}

type PrivateKey interface{}

type Signer interface {
	Public() PublicKey

	Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}

type SignerOpts interface {
	HashFunc() Hash
}

type Decrypter interface {
	Public() PublicKey

	Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}

type DecrypterOpts interface{}

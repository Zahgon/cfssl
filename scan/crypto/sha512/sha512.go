package sha512

import (
	"crypto"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.SHA384, New384)
	crypto.RegisterHash(crypto.SHA512, New)
	crypto.RegisterHash(crypto.SHA512_224, New512_224)
	crypto.RegisterHash(crypto.SHA512_256, New512_256)
}

const (
	Size = 64

	Size224 = 28

	Size256 = 32

	Size384 = 48

	BlockSize = 128
)

const (
	chunk     = 128
	init0     = 0x6a09e667f3bcc908
	init1     = 0xbb67ae8584caa73b
	init2     = 0x3c6ef372fe94f82b
	init3     = 0xa54ff53a5f1d36f1
	init4     = 0x510e527fade682d1
	init5     = 0x9b05688c2b3e6c1f
	init6     = 0x1f83d9abfb41bd6b
	init7     = 0x5be0cd19137e2179
	init0_224 = 0x8c3d37c819544da2
	init1_224 = 0x73e1996689dcd4d6
	init2_224 = 0x1dfab7ae32ff9c82
	init3_224 = 0x679dd514582f9fcf
	init4_224 = 0x0f6d2b697bd44da8
	init5_224 = 0x77e36f7304c48942
	init6_224 = 0x3f9d85a86a1d36c8
	init7_224 = 0x1112e6ad91d692a1
	init0_256 = 0x22312194fc2bf72c
	init1_256 = 0x9f555fa3c84c64c2
	init2_256 = 0x2393b86b6f53b151
	init3_256 = 0x963877195940eabd
	init4_256 = 0x96283ee2a88effe3
	init5_256 = 0xbe5e1e2553863992
	init6_256 = 0x2b0199fc2c85b8aa
	init7_256 = 0x0eb72ddc81c52ca2
	init0_384 = 0xcbbb9d5dc1059ed8
	init1_384 = 0x629a292a367cd507
	init2_384 = 0x9159015a3070dd17
	init3_384 = 0x152fecd8f70e5939
	init4_384 = 0x67332667ffc00b31
	init5_384 = 0x8eb44a8768581511
	init6_384 = 0xdb0c2e0d64f98fa7
	init7_384 = 0x47b5481dbefa4fa4
)

type digest struct {
	h        [8]uint64
	x        [chunk]byte
	nx       int
	len      uint64
	function crypto.Hash
}

func (d *digest) Reset() { _ = "STUB: not implemented"; return }

func New() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func New512_224() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func New512_256() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func New384() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func (d *digest) Size() int { _ = "STUB: not implemented"; return 0 }

func (d *digest) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (d *digest) Write(p []byte) (nn int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (d0 *digest) Sum(in []byte) []byte { _ = "STUB: not implemented"; return nil }

func (d *digest) checkSum() [Size]byte { _ = "STUB: not implemented"; return [Size]byte{} }

func Sum512(data []byte) [Size]byte { _ = "STUB: not implemented"; return [Size]byte{} }

func Sum384(data []byte) (sum384 [Size384]byte) { _ = "STUB: not implemented"; return [Size384]byte{} }

func Sum512_224(data []byte) (sum224 [Size224]byte) {
	_ = "STUB: not implemented"
	return [Size224]byte{}
}

func Sum512_256(data []byte) (sum256 [Size256]byte) {
	_ = "STUB: not implemented"
	return [Size256]byte{}
}

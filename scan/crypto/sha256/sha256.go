package sha256

import (
	"crypto"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.SHA224, New224)
	crypto.RegisterHash(crypto.SHA256, New)
}

const Size = 32

const Size224 = 28

const BlockSize = 64

const (
	chunk     = 64
	init0     = 0x6A09E667
	init1     = 0xBB67AE85
	init2     = 0x3C6EF372
	init3     = 0xA54FF53A
	init4     = 0x510E527F
	init5     = 0x9B05688C
	init6     = 0x1F83D9AB
	init7     = 0x5BE0CD19
	init0_224 = 0xC1059ED8
	init1_224 = 0x367CD507
	init2_224 = 0x3070DD17
	init3_224 = 0xF70E5939
	init4_224 = 0xFFC00B31
	init5_224 = 0x68581511
	init6_224 = 0x64F98FA7
	init7_224 = 0xBEFA4FA4
)

type digest struct {
	h     [8]uint32
	x     [chunk]byte
	nx    int
	len   uint64
	is224 bool
}

func (d *digest) Reset() { _ = "STUB: not implemented"; return }

func New() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func New224() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func (d *digest) Size() int { _ = "STUB: not implemented"; return 0 }

func (d *digest) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (d *digest) Write(p []byte) (nn int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (d0 *digest) Sum(in []byte) []byte { _ = "STUB: not implemented"; return nil }

func (d *digest) checkSum() [Size]byte { _ = "STUB: not implemented"; return [Size]byte{} }

func Sum256(data []byte) [Size]byte { _ = "STUB: not implemented"; return [Size]byte{} }

func Sum224(data []byte) (sum224 [Size224]byte) { _ = "STUB: not implemented"; return [Size224]byte{} }

//go:generate go run gen.go -full -output md5block.go

package md5

import (
	"crypto"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.MD5, New)
}

const Size = 16

const BlockSize = 64

const (
	chunk = 64
	init0 = 0x67452301
	init1 = 0xEFCDAB89
	init2 = 0x98BADCFE
	init3 = 0x10325476
)

type digest struct {
	s   [4]uint32
	x   [chunk]byte
	nx  int
	len uint64
}

func (d *digest) Reset() { _ = "STUB: not implemented"; return }

func New() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func (d *digest) Size() int { _ = "STUB: not implemented"; return 0 }

func (d *digest) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (d *digest) Write(p []byte) (nn int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (d0 *digest) Sum(in []byte) []byte { _ = "STUB: not implemented"; return nil }

func (d *digest) checkSum() [Size]byte { _ = "STUB: not implemented"; return [Size]byte{} }

func Sum(data []byte) [Size]byte { _ = "STUB: not implemented"; return [Size]byte{} }

package md5

import (
	"runtime"
	"unsafe"
)

const x86 = runtime.GOARCH == "amd64" || runtime.GOARCH == "386"

var littleEndian bool

func init() {
	x := uint32(0x04030201)
	y := [4]byte{0x1, 0x2, 0x3, 0x4}
	littleEndian = *(*[4]byte)(unsafe.Pointer(&x)) == y
}

func blockGeneric(dig *digest, p []byte) { _ = "STUB: not implemented"; return }

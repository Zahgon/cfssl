//go:build !amd64 && !amd64p32 && !386 && !arm
// +build !amd64,!amd64p32,!386,!arm

package md5

var block = blockGeneric

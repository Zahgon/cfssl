//go:build amd64 || amd64p32 || 386 || arm
// +build amd64 amd64p32 386 arm

package md5

//go:noescape

func block(dig *digest, p []byte)

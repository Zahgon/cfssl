//go:build 386 || amd64
// +build 386 amd64

package sha256

//go:noescape

func block(dig *digest, p []byte)

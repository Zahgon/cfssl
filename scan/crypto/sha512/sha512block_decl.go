//go:build amd64
// +build amd64

package sha512

//go:noescape

func block(dig *digest, p []byte)

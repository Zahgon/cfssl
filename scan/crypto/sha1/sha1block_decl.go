//go:build amd64 || amd64p32 || arm || 386
// +build amd64 amd64p32 arm 386

package sha1

//go:noescape

func block(dig *digest, p []byte)

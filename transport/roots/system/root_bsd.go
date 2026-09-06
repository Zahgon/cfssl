//go:build dragonfly || freebsd || netbsd || openbsd
// +build dragonfly freebsd netbsd openbsd

package system

var certFiles = []string{
	"/usr/local/share/certs/ca-root-nss.crt",
	"/etc/ssl/cert.pem",
	"/etc/openssl/certs/ca-certificates.crt",
}

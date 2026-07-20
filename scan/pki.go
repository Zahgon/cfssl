package scan

import (
	"crypto/x509"
	"time"

	"github.com/cloudflare/cfssl/scan/crypto/tls"
)

var PKI = &Family{
	Description: "Scans for the Public Key Infrastructure",
	Scanners: map[string]*Scanner{
		"ChainExpiration": {
			"Host's chain hasn't expired and won't expire in the next 30 days",
			chainExpiration,
		},
		"ChainValidation": {
			"All certificates in host's chain are valid",
			chainValidation,
		},
		"MultipleCerts": {
			"Host serves same certificate chain across all IPs",
			multipleCerts,
		},
	},
}

func getChain(addr string, config *tls.Config) (chain []*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type expiration time.Time

func (e expiration) String() string { _ = "STUB: not implemented"; return "" }

func chainExpiration(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

func chainValidation(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

func multipleCerts(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

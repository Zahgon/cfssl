//go:build plan9
// +build plan9

package system

import (
	"crypto/x509"
)

var certFiles = []string{
	"/sys/lib/tls/ca.pem",
}

func initSystemRoots() (roots []*x509.Certificate) { _ = "STUB: not implemented"; return nil }

//go:generate go run root_darwin_arm_gen.go -output root_darwin_armx.go

package system

import (
	"crypto/x509"
)

func execSecurityRoots() ([]*x509.Certificate, error) { _ = "STUB: not implemented"; return nil, nil }

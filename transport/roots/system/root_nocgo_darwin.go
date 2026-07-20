//go:build !cgo
// +build !cgo

package system

import "crypto/x509"

func initSystemRoots() []*x509.Certificate { _ = "STUB: not implemented"; return nil }

//go:build dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris
// +build dragonfly freebsd linux nacl netbsd openbsd solaris

package system

import (
	"crypto/x509"
)

var certDirectories = []string{
	"/system/etc/security/cacerts",
}

func initSystemRoots() []*x509.Certificate { _ = "STUB: not implemented"; return nil }

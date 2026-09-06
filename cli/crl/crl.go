package crl

import (
	"github.com/cloudflare/cfssl/cli"
)

var crlUsageText = `cfssl crl -- generate a new Certificate Revocation List from Database

Usage of crl:
        cfssl crl

Flags:
`
var crlFlags = []string{"db-config", "ca", "ca-key", "expiry"}

func generateCRL(c cli.Config) (crlBytes []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func crlMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: crlUsageText, Flags: crlFlags, Main: crlMain}

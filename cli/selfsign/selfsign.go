package selfsign

import (
	"github.com/cloudflare/cfssl/cli"
)

var selfSignUsageText = `cfssl selfsign -- generate a new self-signed key and signed certificate

Usage of gencert:
        cfssl selfsign HOSTNAME CSRJSON

WARNING: this should ONLY be used for testing. This should never be
used in production.

WARNING: self-signed certificates are insecure; they do not provide
the authentication required for secure systems. Use these at your own
risk.

Arguments:
        HOSTNAME:   Hostname for the cert
        CSRJSON:    JSON file containing the request, use '-' for reading JSON from stdin

Flags:
`

var selfSignFlags = []string{"config"}

func selfSignMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: selfSignUsageText, Flags: selfSignFlags, Main: selfSignMain}

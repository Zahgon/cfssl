package gencrl

import (
	"github.com/cloudflare/cfssl/cli"
)

var gencrlUsageText = `cfssl gencrl -- generate a new Certificate Revocation List

Usage of gencrl:
        cfssl gencrl INPUTFILE CERT KEY TIME

Arguments:
        INPUTFILE:               Text file with one serial number per line, use '-' for reading text from stdin
        CERT:                    The certificate that is signing this CRL, use '-' for reading text from stdin
        KEY:                     The private key of the certificate that is signing the CRL, use '-' for reading text from stdin
        TIME (OPTIONAL):         The desired expiration from now, in seconds

Flags:
`
var gencrlFlags = []string{}

func gencrlMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: gencrlUsageText, Flags: gencrlFlags, Main: gencrlMain}

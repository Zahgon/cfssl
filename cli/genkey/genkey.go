package genkey

import (
	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/csr"
)

var genkeyUsageText = `cfssl genkey -- generate a new key and CSR

Usage of genkey:
        cfssl genkey CSRJSON

Arguments:
        CSRJSON:    JSON file containing the request, use '-' for reading JSON from stdin

Flags:
`

var genkeyFlags = []string{"initca", "config"}

func genkeyMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

func Validator(req *csr.CertificateRequest) error { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: genkeyUsageText, Flags: genkeyFlags, Main: genkeyMain}

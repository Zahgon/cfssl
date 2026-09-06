package gencsr

import (
	"github.com/cloudflare/cfssl/cli"
)

var gencsrUsageText = `cfssl gencsr -- generate a csr from a private key with existing CSR json specification or certificate

Usage of gencsr:
        cfssl gencsr -key private_key_file [-host hostname_override] CSRJSON
        cfssl gencsr -key private_key_file [-host hostname_override] -cert certificate_file

Arguments:
        CSRJSON:    JSON file containing the request, use '-' for reading JSON from stdin

Flags:
`

var gencsrFlags = []string{"key", "cert"}

func gencsrMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: gencsrUsageText, Flags: gencsrFlags, Main: gencsrMain}

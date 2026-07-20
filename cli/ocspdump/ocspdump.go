package ocspdump

import (
	"github.com/cloudflare/cfssl/cli"
)

var ocspdumpUsageText = `cfssl ocspdump -- generates a series of concatenated OCSP responses
for use with ocspserve from all OCSP responses in the cert db

Usage of ocspdump:
        cfssl ocspdump -db-config db-config

Flags:
`

var ocspdumpFlags = []string{"db-config"}

func ocspdumpMain(args []string, c cli.Config) error { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: ocspdumpUsageText, Flags: ocspdumpFlags, Main: ocspdumpMain}

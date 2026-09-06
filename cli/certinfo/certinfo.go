package certinfo

import (
	"github.com/cloudflare/cfssl/cli"
)

var dataUsageText = `cfssl certinfo -- output certinfo about the given cert

Usage of certinfo:
	- Data from local certificate files
        cfssl certinfo -cert file
	- Data from local CSR file
        cfssl certinfo -csr file
	- Data from certificate from remote server.
        cfssl certinfo -domain domain_name
	- Data from CA storage
        cfssl certinfo -sn serial (requires -db-config and -aki)

Flags:
`

var certinfoFlags = []string{"aki", "cert", "csr", "db-config", "domain", "serial"}

func certinfoMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: dataUsageText, Flags: certinfoFlags, Main: certinfoMain}

package bundle

import (
	"github.com/cloudflare/cfssl/cli"
)

var bundlerUsageText = `cfssl bundle -- create a certificate bundle that contains the client cert

Usage of bundle:
	- Bundle local certificate files
        cfssl bundle -cert file [-ca-bundle file] [-int-bundle file] [-int-dir dir] [-metadata file] [-key keyfile] [-flavor optimal|ubiquitous|force] [-password password]
	- Bundle certificate from remote server.
        cfssl bundle -domain domain_name [-ip ip_address] [-ca-bundle file] [-int-bundle file] [-int-dir dir] [-metadata file]

Flags:
`

var bundlerFlags = []string{"cert", "key", "ca-bundle", "int-bundle", "flavor", "int-dir", "metadata", "domain", "ip", "password"}

func bundlerMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: bundlerUsageText, Flags: bundlerFlags, Main: bundlerMain}

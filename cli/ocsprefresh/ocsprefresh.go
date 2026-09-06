package ocsprefresh

import (
	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/ocsp"
)

var ocsprefreshUsageText = `cfssl ocsprefresh -- refreshes the ocsp_responses table
with new OCSP responses for all known unexpired certificates

Usage of ocsprefresh:
        cfssl ocsprefresh -db-config db-config -ca cert -responder cert -responder-key key [-interval 96h]

Flags:
`

var ocsprefreshFlags = []string{"ca", "responder", "responder-key", "db-config", "interval"}

func ocsprefreshMain(args []string, c cli.Config) error { _ = "STUB: not implemented"; return nil }

func SignerFromConfig(c cli.Config) (ocsp.Signer, error) {
	_ = "STUB: not implemented"
	return *new(ocsp.Signer), nil
}

var Command = &cli.Command{UsageText: ocsprefreshUsageText, Flags: ocsprefreshFlags, Main: ocsprefreshMain}

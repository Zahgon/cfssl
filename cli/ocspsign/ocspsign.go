package ocspsign

import (
	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/ocsp"
)

var ocspSignerUsageText = `cfssl ocspsign -- signs an OCSP response for a given CA, cert, and status.
Returns a base64-encoded OCSP response.

Usage of ocspsign:
        cfssl ocspsign -ca cert -responder cert -responder-key key -cert cert [-status status] [-reason code] [-revoked-at YYYY-MM-DD] [-interval 96h]

Flags:
`

var ocspSignerFlags = []string{"ca", "responder", "responder-key", "reason", "status", "revoked-at", "interval"}

func ocspSignerMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

func SignerFromConfig(c cli.Config) (ocsp.Signer, error) {
	_ = "STUB: not implemented"
	return *new(ocsp.Signer), nil
}

var Command = &cli.Command{UsageText: ocspSignerUsageText, Flags: ocspSignerFlags, Main: ocspSignerMain}

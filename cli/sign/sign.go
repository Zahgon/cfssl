package sign

import (
	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/signer"

	"github.com/jmoiron/sqlx"
)

var signerUsageText = `cfssl sign -- signs a client cert with a host name by a given CA and CA key

Usage of sign:
        cfssl sign -ca cert -ca-key key [mutual-tls-cert cert] [mutual-tls-key key] [-config config] [-profile profile] [-hostname hostname] [-db-config db-config] CSR [SUBJECT]
        cfssl sign -remote remote_host [mutual-tls-cert cert] [mutual-tls-key key] [-config config] [-profile profile] [-label label] [-hostname hostname] CSR [SUBJECT]

Arguments:
        CSR:        PEM file for certificate request, use '-' for reading PEM from stdin.

Note: CSR can also be supplied via flag values; flag value will take precedence over the argument.

SUBJECT is an optional file containing subject information to use for the certificate instead of the subject information in the CSR.

Flags:
`

var signerFlags = []string{"hostname", "csr", "ca", "ca-key", "config", "profile", "label", "remote",
	"mutual-tls-cert", "mutual-tls-key", "db-config"}

func SignerFromConfigAndDB(c cli.Config, db *sqlx.DB) (signer.Signer, error) {
	_ = "STUB: not implemented"
	return *new(signer.Signer), nil
}

func SignerFromConfig(c cli.Config) (s signer.Signer, err error) {
	_ = "STUB: not implemented"
	return *new(signer.Signer), nil
}

func signerMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: signerUsageText, Flags: signerFlags, Main: signerMain}

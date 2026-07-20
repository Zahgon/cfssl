package gencert

import (
	"github.com/cloudflare/cfssl/cli"
)

var gencertUsageText = `cfssl gencert -- generate a new key and signed certificate

Usage of gencert:
    Generate a new key and cert from CSR:
        cfssl gencert -initca CSRJSON
        cfssl gencert -ca cert -ca-key key [-config config] [-profile profile] [-hostname hostname] CSRJSON
        cfssl gencert -remote remote_host [-config config] [-profile profile] [-label label] [-hostname hostname] CSRJSON

    Re-generate a CA cert with the CA key and CSR:
        cfssl gencert -initca -ca-key key CSRJSON

    Re-generate a CA cert with the CA key and certificate:
        cfssl gencert -renewca -ca cert -ca-key key

Arguments:
        CSRJSON:    JSON file containing the request, use '-' for reading JSON from stdin

Flags:
`

var gencertFlags = []string{"initca", "remote", "ca", "ca-key", "config", "cn", "hostname", "profile", "label"}

func gencertMain(args []string, c cli.Config) error { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: gencertUsageText, Flags: gencertFlags, Main: gencertMain}

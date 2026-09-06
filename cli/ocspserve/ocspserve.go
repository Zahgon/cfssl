package ocspserve

import (
	"github.com/cloudflare/cfssl/cli"
)

var ocspServerUsageText = `cfssl ocspserve -- set up an HTTP server that handles OCSP requests from either a file or directly from a database (see RFC 5019)

  Usage of ocspserve:
          cfssl ocspserve [-address address] [-port port] [-responses file] [-db-config db-config]

  Flags:
  `

var ocspServerFlags = []string{"address", "port", "responses", "db-config"}

func ocspServerMain(args []string, c cli.Config) error { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: ocspServerUsageText, Flags: ocspServerFlags, Main: ocspServerMain}

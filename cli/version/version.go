package version

import (
	"github.com/cloudflare/cfssl/cli"
)

var (
	version = "dev"
)

var versionUsageText = `cfssl version -- print out the version of CF SSL

Usage of version:
	cfssl version
`

func FormatVersion() string { _ = "STUB: not implemented"; return "" }

func versionMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: versionUsageText, Flags: nil, Main: versionMain}

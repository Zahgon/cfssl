package printdefaults

import (
	"github.com/cloudflare/cfssl/cli"
)

var printDefaultsUsage = `cfssl print-defaults -- print default configurations that can be used as a template

Usage of print-defaults:
        cfssl print-defaults TYPE

If "list" is used as the TYPE, the list of supported types will be printed.
`

func printAvailable() { _ = "STUB: not implemented"; return }

func printDefaults(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{
	UsageText: printDefaultsUsage,
	Flags:     []string{},
	Main:      printDefaults,
}

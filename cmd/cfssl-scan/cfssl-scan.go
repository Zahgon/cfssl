package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/scan"
	"github.com/cloudflare/cfssl/config"
)

func main() {

	var scanFlagSet = flag.NewFlagSet("scan", flag.ExitOnError)
	var c cli.Config
	var usageText = `cfssl scan -- scan a host for issues
Usage of scan:
        cfssl scan [-family regexp] [-scanner regexp] [-timeout duration] [-ip IPAddr] [-num-workers num] [-max-hosts num] [-csv hosts.csv] HOST+
        cfssl scan -list

Arguments:
        HOST:    Host(s) to scan (including port)
Flags:
`
	registerFlags(&c, scanFlagSet)

	scanFlagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "\t%s", usageText)
		for _, name := range scan.Command.Flags {
			if f := scanFlagSet.Lookup(name); f != nil {
				printDefaultValue(f)
			}
		}
	}
	args := os.Args[1:]
	scanFlagSet.Parse(args)
	args = scanFlagSet.Args()

	var err error
	c.CFG, err = config.LoadFile(c.ConfigFile)
	if c.ConfigFile != "" && err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config file: %v", err)
	}

	if err := scan.Command.Main(args, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func printDefaultValue(f *flag.Flag) { _ = "STUB: not implemented"; return }

func registerFlags(c *cli.Config, f *flag.FlagSet) { _ = "STUB: not implemented"; return }

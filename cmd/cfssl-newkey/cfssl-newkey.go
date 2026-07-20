package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/genkey"
	"github.com/cloudflare/cfssl/config"
)

func main() {

	var newkeyFlagSet = flag.NewFlagSet("newkey", flag.ExitOnError)
	var c cli.Config

	var usageText = `cfssl-newkey -- generate a new key and CSR

	Usage of genkey:
        	newkey CSRJSON

	Arguments:
        	CSRJSON:    JSON file containing the request, use '-' for reading JSON from stdin

	Flags:
	`

	registerFlags(&c, newkeyFlagSet)

	newkeyFlagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "\t%s", usageText)
		for _, name := range genkey.Command.Flags {
			if f := newkeyFlagSet.Lookup(name); f != nil {
				printDefaultValue(f)
			}
		}
	}
	args := os.Args[1:]
	newkeyFlagSet.Parse(args)
	args = newkeyFlagSet.Args()

	var err error
	c.CFG, err = config.LoadFile(c.ConfigFile)
	if c.ConfigFile != "" && err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config file: %v", err)
	}

	if err := genkey.Command.Main(args, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func printDefaultValue(f *flag.Flag) { _ = "STUB: not implemented"; return }

func registerFlags(c *cli.Config, f *flag.FlagSet) { _ = "STUB: not implemented"; return }

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/bundle"
	"github.com/cloudflare/cfssl/config"
)

func main() {

	var bundleFlagSet = flag.NewFlagSet("bundle", flag.ExitOnError)
	var c cli.Config
	var usageText = `cfssl-bundle -- create a certificate bundle that contains the client cert

	Usage of bundle:
		- Bundle local certificate files
        	bundle -cert file [-ca-bundle file] [-int-bundle file] [-int-dir dir] [-metadata file] [-key keyfile] [-flavor optimal|ubiquitous|force] [-password password]
		- Bundle certificate from remote server.
        	bundle -domain domain_name [-ip ip_address] [-ca-bundle file] [-int-bundle file] [-int-dir dir] [-metadata file]

	Flags:
	`

	registerFlags(&c, bundleFlagSet)

	bundleFlagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "\t%s", usageText)
		for _, name := range bundle.Command.Flags {
			if f := bundleFlagSet.Lookup(name); f != nil {
				printDefaultValue(f)
			}
		}
	}
	args := os.Args[1:]
	bundleFlagSet.Parse(args)
	args = bundleFlagSet.Args()

	var err error
	c.CFG, err = config.LoadFile(c.ConfigFile)
	if c.ConfigFile != "" && err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config file: %v", err)
	}

	if err := bundle.Command.Main(args, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func printDefaultValue(f *flag.Flag) { _ = "STUB: not implemented"; return }

func registerFlags(c *cli.Config, f *flag.FlagSet) { _ = "STUB: not implemented"; return }

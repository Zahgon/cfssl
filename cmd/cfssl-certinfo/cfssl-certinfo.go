package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/certinfo"
	"github.com/cloudflare/cfssl/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var certinfoFlagSet = flag.NewFlagSet("certinfo", flag.ExitOnError)
	var c cli.Config
	registerFlags(&c, certinfoFlagSet)
	var usageText = `cfssl-certinfo -- output certinfo about the given cert

	Usage of certinfo:
		- Data from local certificate files
        	certinfo -cert file
		- Data from certificate from remote server.
        	certinfo -domain domain_name
		- Data from CA storage
        	certinfo -serial serial_number -aki authority_key_id (requires -db-config)

	Flags:
	`

	certinfoFlagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "\t%s", usageText)
		for _, name := range certinfo.Command.Flags {
			if f := certinfoFlagSet.Lookup(name); f != nil {
				printDefaultValue(f)
			}
		}
	}
	args := os.Args[1:]
	certinfoFlagSet.Parse(args)
	args = certinfoFlagSet.Args()

	var err error
	c.CFG, err = config.LoadFile(c.ConfigFile)
	if c.ConfigFile != "" && err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config file: %v", err)
	}

	if err := certinfo.Command.Main(args, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func printDefaultValue(f *flag.Flag) { _ = "STUB: not implemented"; return }

func registerFlags(c *cli.Config, f *flag.FlagSet) { _ = "STUB: not implemented"; return }

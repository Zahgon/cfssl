package main

import (
	"flag"
	"os"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/bundle"
	"github.com/cloudflare/cfssl/cli/certinfo"
	"github.com/cloudflare/cfssl/cli/crl"
	"github.com/cloudflare/cfssl/cli/gencert"
	"github.com/cloudflare/cfssl/cli/gencrl"
	"github.com/cloudflare/cfssl/cli/gencsr"
	"github.com/cloudflare/cfssl/cli/genkey"
	"github.com/cloudflare/cfssl/cli/info"
	"github.com/cloudflare/cfssl/cli/ocspdump"
	"github.com/cloudflare/cfssl/cli/ocsprefresh"
	"github.com/cloudflare/cfssl/cli/ocspserve"
	"github.com/cloudflare/cfssl/cli/ocspsign"
	"github.com/cloudflare/cfssl/cli/revoke"
	"github.com/cloudflare/cfssl/cli/scan"
	"github.com/cloudflare/cfssl/cli/selfsign"
	"github.com/cloudflare/cfssl/cli/serve"
	"github.com/cloudflare/cfssl/cli/sign"
	"github.com/cloudflare/cfssl/cli/version"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	flag.Usage = nil

	cmds := map[string]*cli.Command{
		"bundle":         bundle.Command,
		"certinfo":       certinfo.Command,
		"crl":            crl.Command,
		"sign":           sign.Command,
		"serve":          serve.Command,
		"version":        version.Command,
		"genkey":         genkey.Command,
		"gencert":        gencert.Command,
		"gencsr":         gencsr.Command,
		"gencrl":         gencrl.Command,
		"ocspdump":       ocspdump.Command,
		"ocsprefresh":    ocsprefresh.Command,
		"ocspsign":       ocspsign.Command,
		"ocspserve":      ocspserve.Command,
		"selfsign":       selfsign.Command,
		"scan":           scan.Command,
		"info":           info.Command,
		"print-defaults": printdefaults.Command,
		"revoke":         revoke.Command,
	}

	err := cli.Start(cmds)
	if err != nil {
		os.Exit(1)
	}
}

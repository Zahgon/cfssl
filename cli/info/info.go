package info

import (
	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/info"
)

var infoUsageTxt = `cfssl info -- get info about a remote signer

Usage:

Get info about a remote signer:
cfssl info -remote remote_host [-label label] [-profile profile] [-label label] 

Flags:
`

var infoFlags = []string{"remote", "label", "profile", "config"}

func getInfoFromRemote(c cli.Config) (resp *info.Resp, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getInfoFromConfig(c cli.Config) (resp *info.Resp, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func infoMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{
	UsageText: infoUsageTxt,
	Flags:     infoFlags,
	Main:      infoMain,
}

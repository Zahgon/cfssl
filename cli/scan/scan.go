package scan

import (
	"sync"

	"github.com/cloudflare/cfssl/cli"
)

var scanUsageText = `cfssl scan -- scan a host for issues
Usage of scan:
        cfssl scan [-family regexp] [-scanner regexp] [-timeout duration] [-ip IPAddr] [-num-workers num] [-max-hosts num] [-csv hosts.csv] HOST+
        cfssl scan -list

Arguments:
        HOST:    Host(s) to scan (including port)
Flags:
`
var scanFlags = []string{"list", "family", "scanner", "timeout", "ip", "ca-bundle", "num-workers", "csv", "max-hosts"}

func printJSON(v interface{}) { _ = "STUB: not implemented"; return }

type context struct {
	sync.WaitGroup
	c     cli.Config
	hosts chan string
}

func newContext(c cli.Config, numWorkers int) *context { _ = "STUB: not implemented"; return nil }

func (ctx *context) runWorker() { _ = "STUB: not implemented"; return }

func parseCSV(hosts []string, csvFile string, maxHosts int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanMain(args []string, c cli.Config) (err error) { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: scanUsageText, Flags: scanFlags, Main: scanMain}

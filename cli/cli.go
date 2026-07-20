package cli

import (
	"flag"
)

type Command struct {
	UsageText string

	Flags []string

	Main func(args []string, c Config) error
}

var cmdName string

const usage = `Usage:
Available commands:
`

func printDefaultValue(f *flag.Flag) { _ = "STUB: not implemented"; return }

func PopFirstArgument(args []string) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func Start(cmds map[string]*Command) error { _ = "STUB: not implemented"; return nil }

func ReadStdin(filename string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func PrintCert(key, csrBytes, cert []byte) { _ = "STUB: not implemented"; return }

func PrintOCSPResponse(resp []byte) { _ = "STUB: not implemented"; return }

func PrintCRL(certList []byte) { _ = "STUB: not implemented"; return }

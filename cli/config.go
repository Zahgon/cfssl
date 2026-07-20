package cli

import (
	"flag"
	"time"

	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/signer/universal"
)

type Config struct {
	Hostname          string
	CertFile          string
	CSRFile           string
	CAFile            string
	CAKeyFile         string
	TLSCertFile       string
	TLSKeyFile        string
	MutualTLSCAFile   string
	MutualTLSCNRegex  string
	TLSRemoteCAs      string
	MutualTLSCertFile string
	MutualTLSKeyFile  string
	KeyFile           string
	IntermediatesFile string
	CABundleFile      string
	IntBundleFile     string
	Address           string
	Port              int
	MinTLSVersion     string
	Password          string
	ConfigFile        string
	CFG               *config.Config
	Profile           string
	IsCA              bool
	RenewCA           bool
	IntDir            string
	Flavor            string
	Metadata          string
	Domain            string
	IP                string
	Remote            string
	Label             string
	AuthKey           string
	ResponderFile     string
	ResponderKeyFile  string
	Status            string
	Reason            string
	RevokedAt         string
	Interval          time.Duration
	List              bool
	Family            string
	Timeout           time.Duration
	Scanner           string
	CSVFile           string
	NumWorkers        int
	MaxHosts          int
	Responses         string
	Path              string
	CRL               string
	Usage             string
	PGPPrivate        string
	PGPName           string
	Serial            string
	CNOverride        string
	AKI               string
	DBConfigFile      string
	CRLExpiration     time.Duration
	Disable           string
}

func registerFlags(c *Config, f *flag.FlagSet) { _ = "STUB: not implemented"; return }

func RootFromConfig(c *Config) universal.Root {
	_ = "STUB: not implemented"
	return *new(universal.Root)
}

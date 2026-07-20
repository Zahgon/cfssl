package config

import (
	"crypto"
	"crypto/x509"
	"errors"
	"regexp"

	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/whitelist"

	"github.com/jmoiron/sqlx"
)

type RawMap map[string]map[string]string

var (
	configSection    = regexp.MustCompile("^\\s*\\[\\s*(\\w+)\\s*\\]\\s*$")
	quotedConfigLine = regexp.MustCompile("^\\s*(\\w+)\\s*=\\s*[\"'](.*)[\"']\\s*$")
	configLine       = regexp.MustCompile("^\\s*(\\w+)\\s*=\\s*(.*)\\s*$")
	commentLine      = regexp.MustCompile("^#.*$")
	blankLine        = regexp.MustCompile("^\\s*$")

	defaultSection = "default"
)

func ParseToRawMap(fileName string) (cfg RawMap, err error) {
	_ = "STUB: not implemented"
	return *new(RawMap), nil
}

func (c *RawMap) SectionInConfig(section string) bool { _ = "STUB: not implemented"; return false }

type Root struct {
	PrivateKey  crypto.Signer
	Certificate *x509.Certificate
	Config      *config.Signing
	ACL         whitelist.NetACL
	DB          *sqlx.DB
}

func LoadRoot(cfg map[string]string) (*Root, error) { _ = "STUB: not implemented"; return nil, nil }

func parsePrivateKeySpec(spec string, cfg map[string]string) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

func parseACL(nets string) (whitelist.NetACL, error) {
	_ = "STUB: not implemented"
	return *new(whitelist.NetACL), nil
}

type RootList map[string]*Root

var (
	ErrMissingPrivateKey = errors.New("config: root is missing private key spec")

	ErrMissingCertificatePath = errors.New("config: root is missing certificate path")

	ErrMissingConfigPath = errors.New("config: root is missing configuration file path")

	ErrInvalidConfig = errors.New("config: invalid configuration")

	ErrUnsupportedScheme = errors.New("config: unsupported private key scheme")
)

func Parse(filename string) (RootList, error) {
	_ = "STUB: not implemented"
	return *new(RootList), nil
}

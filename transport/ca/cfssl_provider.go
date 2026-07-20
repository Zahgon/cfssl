package ca

import (
	"errors"
	"net"

	"github.com/cloudflare/cfssl/api/client"
	"github.com/cloudflare/cfssl/auth"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/transport/core"
)

type authError struct {
	authType string
}

func (err *authError) Error() string { _ = "STUB: not implemented"; return "" }

var authTypes = map[string]func(config.AuthKey, []byte) (auth.Provider, error){
	"standard": newStandardProvider,
}

func newStandardProvider(ak config.AuthKey, ad []byte) (auth.Provider, error) {
	_ = "STUB: not implemented"
	return *new(auth.Provider), nil
}

func newProvider(ak config.AuthKey, ad []byte) (auth.Provider, error) {
	_ = "STUB: not implemented"
	return *new(auth.Provider), nil
}

var ErrNoAuth = errors.New("transport: authentication is required for non-local remotes")

var v4Loopback = net.IPNet{
	IP:   net.IP{127, 0, 0, 0},
	Mask: net.IPv4Mask(255, 0, 0, 0),
}

func ipIsLocal(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func (cap *CFSSL) validateAuth() error { _ = "STUB: not implemented"; return nil }

var cfsslConfigDirs = []string{
	"/usr/local/cfssl",
	"/etc/cfssl",
	"/state/etc/cfssl",
}

func findLabel(label string) *config.Config { _ = "STUB: not implemented"; return nil }

func getProfile(cfg *config.Config, profileName string) (*config.SigningProfile, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (cap *CFSSL) loadAuth() error { _ = "STUB: not implemented"; return nil }

func getRemote(cfg *config.Config, profile *config.SigningProfile) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (cap *CFSSL) setRemoteAndAuth() error { _ = "STUB: not implemented"; return nil }

type CFSSL struct {
	remote        client.Remote
	provider      auth.Provider
	Profile       string
	Label         string
	DefaultRemote client.Remote
	DefaultAuth   config.AuthKey
}

func (cap *CFSSL) SignCSR(csrPEM []byte) (cert []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cap *CFSSL) CACertificate() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewCFSSLProvider(id *core.Identity, defaultRemote client.Remote) (*CFSSL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

package client

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cfssl/auth"
	"github.com/cloudflare/cfssl/info"
)

type Strategy int

const (
	StrategyInvalid = iota

	StrategyOrderedList
)

var strategyStrings = map[string]Strategy{
	"ordered_list": StrategyOrderedList,
}

func StrategyFromString(s string) Strategy { _ = "STUB: not implemented"; return *new(Strategy) }

func NewGroup(remotes []string, tlsConfig *tls.Config, strategy Strategy) (Remote, error) {
	_ = "STUB: not implemented"
	return *new(Remote), nil
}

type orderedListGroup struct {
	remotes []*server
}

func (g *orderedListGroup) Hosts() []string { _ = "STUB: not implemented"; return nil }

func (g *orderedListGroup) SetRequestTimeout(timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (g *orderedListGroup) SetProxy(proxy func(*http.Request) (*url.URL, error)) {
	_ = "STUB: not implemented"
	return
}

func newOrdererdListGroup(remotes []*server) (Remote, error) {
	_ = "STUB: not implemented"
	return *new(Remote), nil
}

func (g *orderedListGroup) AuthSign(req, id []byte, provider auth.Provider) (resp []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *orderedListGroup) Sign(jsonData []byte) (resp []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *orderedListGroup) Info(jsonData []byte) (resp *info.Resp, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *orderedListGroup) SetReqModifier(mod func(*http.Request, []byte)) {
	_ = "STUB: not implemented"
	return
}

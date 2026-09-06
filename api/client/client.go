package client

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cfssl/api"
	"github.com/cloudflare/cfssl/auth"
	"github.com/cloudflare/cfssl/info"
)

type server struct {
	URL            string
	TLSConfig      *tls.Config
	reqModifier    func(*http.Request, []byte)
	RequestTimeout time.Duration
	proxy          func(*http.Request) (*url.URL, error)
}

type Remote interface {
	AuthSign(req, id []byte, provider auth.Provider) ([]byte, error)
	Sign(jsonData []byte) ([]byte, error)
	Info(jsonData []byte) (*info.Resp, error)
	Hosts() []string
	SetReqModifier(func(*http.Request, []byte))
	SetRequestTimeout(d time.Duration)
	SetProxy(func(*http.Request) (*url.URL, error))
}

func NewServer(addr string) Remote { _ = "STUB: not implemented"; return *new(Remote) }

func NewServerTLS(addr string, tlsConfig *tls.Config) Remote {
	_ = "STUB: not implemented"
	return *new(Remote)
}

func (srv *server) Hosts() []string { _ = "STUB: not implemented"; return nil }

func (srv *server) SetReqModifier(mod func(*http.Request, []byte)) {
	_ = "STUB: not implemented"
	return
}

func (srv *server) SetRequestTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

func (srv *server) SetProxy(proxy func(*http.Request) (*url.URL, error)) {
	_ = "STUB: not implemented"
	return
}

func newServer(u *url.URL, tlsConfig *tls.Config) *server { _ = "STUB: not implemented"; return nil }

func (srv *server) getURL(endpoint string) string { _ = "STUB: not implemented"; return "" }

func (srv *server) createTransport() *http.Transport { _ = "STUB: not implemented"; return nil }

func (srv *server) post(url string, jsonData []byte) (*api.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (srv *server) AuthSign(req, id []byte, provider auth.Provider) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (srv *server) AuthInfo(req, id []byte, provider auth.Provider) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (srv *server) authReq(req, ID []byte, provider auth.Provider, target string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (srv *server) Sign(jsonData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (srv *server) Info(jsonData []byte) (*info.Resp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (srv *server) getResultMap(jsonData []byte, target string) (result map[string]interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (srv *server) request(jsonData []byte, target string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type AuthRemote struct {
	Remote
	provider auth.Provider
}

func NewAuthServer(addr string, tlsConfig *tls.Config, provider auth.Provider) *AuthRemote {
	_ = "STUB: not implemented"
	return nil
}

func (ar *AuthRemote) Sign(req []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func normalizeURL(addr string) (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

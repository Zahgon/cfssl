package tls

import (
	"crypto"
	"net"
)

func Server(conn net.Conn, config *Config) *Conn { _ = "STUB: not implemented"; return nil }

func Client(conn net.Conn, config *Config) *Conn { _ = "STUB: not implemented"; return nil }

type listener struct {
	net.Listener
	config *Config
}

func (l *listener) Accept() (c net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func NewListener(inner net.Listener, config *Config) net.Listener {
	_ = "STUB: not implemented"
	return *new(net.Listener)
}

func Listen(network, laddr string, config *Config) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

type timeoutError struct{}

func (timeoutError) Error() string   { _ = "STUB: not implemented"; return "" }
func (timeoutError) Timeout() bool   { _ = "STUB: not implemented"; return false }
func (timeoutError) Temporary() bool { _ = "STUB: not implemented"; return false }

func DialWithDialer(dialer *net.Dialer, network, addr string, config *Config) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Dial(network, addr string, config *Config) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadX509KeyPair(certFile, keyFile string) (Certificate, error) {
	_ = "STUB: not implemented"
	return *new(Certificate), nil
}

func X509KeyPair(certPEMBlock, keyPEMBlock []byte) (Certificate, error) {
	_ = "STUB: not implemented"
	return *new(Certificate), nil
}

func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

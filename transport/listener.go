package transport

import (
	"crypto/tls"
	"net"
	"time"
)

type Listener struct {
	*Transport
	net.Listener
}

func Listen(address string, tr *Transport) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tr *Transport) getConfig() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

var PollInterval = 30 * time.Second

func pollWait(target time.Time) { _ = "STUB: not implemented"; return }

func (l *Listener) AutoUpdate(certUpdates chan<- time.Time, errChan chan<- error) {
	_ = "STUB: not implemented"
	return
}

package tls

import (
	"net"
)

type clientHandshakeState struct {
	c            *Conn
	serverHello  *serverHelloMsg
	hello        *clientHelloMsg
	suite        *cipherSuite
	finishedHash finishedHash
	masterSecret []byte
	session      *ClientSessionState
}

func (c *Conn) clientHandshake() error { _ = "STUB: not implemented"; return nil }

func (hs *clientHandshakeState) doFullHandshake() error { _ = "STUB: not implemented"; return nil }

func (hs *clientHandshakeState) establishKeys() error { _ = "STUB: not implemented"; return nil }

func (hs *clientHandshakeState) serverResumedSession() bool {
	_ = "STUB: not implemented"
	return false
}

func (hs *clientHandshakeState) processServerHello() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (hs *clientHandshakeState) readFinished(out []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (hs *clientHandshakeState) readSessionTicket() error { _ = "STUB: not implemented"; return nil }

func (hs *clientHandshakeState) sendFinished(out []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func clientSessionCacheKey(serverAddr net.Addr, config *Config) string {
	_ = "STUB: not implemented"
	return ""
}

func mutualProtocol(protos, preferenceProtos []string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

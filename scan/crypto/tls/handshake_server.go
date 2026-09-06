package tls

import (
	"crypto"
)

type serverHandshakeState struct {
	c               *Conn
	clientHello     *clientHelloMsg
	hello           *serverHelloMsg
	suite           *cipherSuite
	ellipticOk      bool
	ecdsaOk         bool
	rsaDecryptOk    bool
	rsaSignOk       bool
	sessionState    *sessionState
	finishedHash    finishedHash
	masterSecret    []byte
	certsFromClient [][]byte
	cert            *Certificate
}

func (c *Conn) serverHandshake() error { _ = "STUB: not implemented"; return nil }

func (hs *serverHandshakeState) readClientHello() (isResume bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (hs *serverHandshakeState) checkForResumption() bool { _ = "STUB: not implemented"; return false }

func (hs *serverHandshakeState) doResumeHandshake() error { _ = "STUB: not implemented"; return nil }

func (hs *serverHandshakeState) doFullHandshake() error { _ = "STUB: not implemented"; return nil }

func (hs *serverHandshakeState) establishKeys() error { _ = "STUB: not implemented"; return nil }

func (hs *serverHandshakeState) readFinished(out []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (hs *serverHandshakeState) sendSessionTicket() error { _ = "STUB: not implemented"; return nil }

func (hs *serverHandshakeState) sendFinished(out []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (hs *serverHandshakeState) processCertsFromClient(certificates [][]byte) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

func (hs *serverHandshakeState) setCipherSuite(id uint16, supportedCipherSuites []uint16, version uint16) bool {
	_ = "STUB: not implemented"
	return false
}

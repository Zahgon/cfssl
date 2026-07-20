package tls

type sessionState struct {
	vers         uint16
	cipherSuite  uint16
	masterSecret []byte
	certificates [][]byte

	usedOldKey bool
}

func (s *sessionState) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (s *sessionState) marshal() []byte { _ = "STUB: not implemented"; return nil }

func (s *sessionState) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

func (c *Conn) encryptTicket(state *sessionState) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) decryptTicket(encrypted []byte) (*sessionState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

package tls

func (c *Conn) SayHello(newSigAls []SignatureAndHash) (cipherID, curveType uint16, curveID CurveID, version uint16, certs [][]byte, err error) {
	_ = "STUB: not implemented"
	return 0, 0, *new(CurveID), 0, nil, nil
}

func (c *Conn) sayHello(hello *clientHelloMsg) (serverHello *serverHelloMsg, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) exchangeKeys() (serverKeyExchange *serverKeyExchangeMsg, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

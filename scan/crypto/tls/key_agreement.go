package tls

import (
	"crypto"
	"crypto/elliptic"
	"crypto/x509"
	"errors"
	"math/big"
)

var errClientKeyExchange = errors.New("tls: invalid ClientKeyExchange message")
var errServerKeyExchange = errors.New("tls: invalid ServerKeyExchange message")

type rsaKeyAgreement struct{}

func (ka rsaKeyAgreement) generateServerKeyExchange(config *Config, cert *Certificate, clientHello *clientHelloMsg, hello *serverHelloMsg) (*serverKeyExchangeMsg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ka rsaKeyAgreement) processClientKeyExchange(config *Config, cert *Certificate, ckx *clientKeyExchangeMsg, version uint16) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ka rsaKeyAgreement) processServerKeyExchange(config *Config, clientHello *clientHelloMsg, serverHello *serverHelloMsg, cert *x509.Certificate, skx *serverKeyExchangeMsg) error {
	_ = "STUB: not implemented"
	return nil
}

func (ka rsaKeyAgreement) generateClientKeyExchange(config *Config, clientHello *clientHelloMsg, cert *x509.Certificate) ([]byte, *clientKeyExchangeMsg, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func sha1Hash(slices [][]byte) []byte { _ = "STUB: not implemented"; return nil }

func md5SHA1Hash(slices [][]byte) []byte { _ = "STUB: not implemented"; return nil }

func hashForServerKeyExchange(sigAndHash signatureAndHash, version uint16, slices ...[]byte) ([]byte, crypto.Hash, error) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.Hash), nil
}

func pickTLS12HashForSignature(sigType uint8, clientList []signatureAndHash) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func curveForCurveID(id CurveID) (elliptic.Curve, bool) {
	_ = "STUB: not implemented"
	return *new(elliptic.Curve), false
}

type ecdheKeyAgreement struct {
	version    uint16
	sigType    uint8
	privateKey []byte
	curve      elliptic.Curve
	x, y       *big.Int
}

func (ka *ecdheKeyAgreement) generateServerKeyExchange(config *Config, cert *Certificate, clientHello *clientHelloMsg, hello *serverHelloMsg) (*serverKeyExchangeMsg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ka *ecdheKeyAgreement) processClientKeyExchange(config *Config, cert *Certificate, ckx *clientKeyExchangeMsg, version uint16) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ka *ecdheKeyAgreement) processServerKeyExchange(config *Config, clientHello *clientHelloMsg, serverHello *serverHelloMsg, cert *x509.Certificate, skx *serverKeyExchangeMsg) error {
	_ = "STUB: not implemented"
	return nil
}

func (ka *ecdheKeyAgreement) generateClientKeyExchange(config *Config, clientHello *clientHelloMsg, cert *x509.Certificate) ([]byte, *clientKeyExchangeMsg, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

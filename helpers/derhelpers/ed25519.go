package derhelpers

import (
	"crypto"
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
)

var errEd25519WrongID = errors.New("incorrect object identifier")
var errEd25519WrongKeyType = errors.New("incorrect key type")

var ed25519OID = asn1.ObjectIdentifier{1, 3, 101, 112}

type subjectPublicKeyInfo struct {
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}

func MarshalEd25519PublicKey(pk crypto.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseEd25519PublicKey(der []byte) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

type oneAsymmetricKey struct {
	Version    int
	Algorithm  pkix.AlgorithmIdentifier
	PrivateKey []byte
}

type curvePrivateKey []byte

func MarshalEd25519PrivateKey(sk crypto.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseEd25519PrivateKey(der []byte) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

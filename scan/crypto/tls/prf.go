package tls

import (
	"crypto"
	"hash"
)

func splitPreMasterSecret(secret []byte) (s1, s2 []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pHash(result, secret, seed []byte, hash func() hash.Hash) { _ = "STUB: not implemented"; return }

func prf10(result, secret, label, seed []byte) { _ = "STUB: not implemented"; return }

func prf12(hashFunc func() hash.Hash) func(result, secret, label, seed []byte) {
	_ = "STUB: not implemented"
	return nil
}

func prf30(result, secret, label, seed []byte) { _ = "STUB: not implemented"; return }

const (
	tlsRandomLength      = 32
	masterSecretLength   = 48
	finishedVerifyLength = 12
)

var masterSecretLabel = []byte("master secret")
var keyExpansionLabel = []byte("key expansion")
var clientFinishedLabel = []byte("client finished")
var serverFinishedLabel = []byte("server finished")

func prfAndHashForVersion(version uint16, suite *cipherSuite) (func(result, secret, label, seed []byte), crypto.Hash) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.Hash)
}

func prfForVersion(version uint16, suite *cipherSuite) func(result, secret, label, seed []byte) {
	_ = "STUB: not implemented"
	return nil
}

func masterFromPreMasterSecret(version uint16, suite *cipherSuite, preMasterSecret, clientRandom, serverRandom []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func keysFromMasterSecret(version uint16, suite *cipherSuite, masterSecret, clientRandom, serverRandom []byte, macLen, keyLen, ivLen int) (clientMAC, serverMAC, clientKey, serverKey, clientIV, serverIV []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil, nil
}

func lookupTLSHash(hash uint8) (crypto.Hash, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Hash), nil
}

func newFinishedHash(version uint16, cipherSuite *cipherSuite) finishedHash {
	_ = "STUB: not implemented"
	return *new(finishedHash)
}

type finishedHash struct {
	client hash.Hash
	server hash.Hash

	clientMD5 hash.Hash
	serverMD5 hash.Hash

	buffer []byte

	version uint16
	prf     func(result, secret, label, seed []byte)
}

func (h *finishedHash) Write(msg []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h finishedHash) Sum() []byte { _ = "STUB: not implemented"; return nil }

func finishedSum30(md5, sha1 hash.Hash, masterSecret []byte, magic []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

var ssl3ClientFinishedMagic = [4]byte{0x43, 0x4c, 0x4e, 0x54}
var ssl3ServerFinishedMagic = [4]byte{0x53, 0x52, 0x56, 0x52}

func (h finishedHash) clientSum(masterSecret []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h finishedHash) serverSum(masterSecret []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h finishedHash) selectClientCertSignatureAlgorithm(serverList []signatureAndHash, sigType uint8) (signatureAndHash, error) {
	_ = "STUB: not implemented"
	return *new(signatureAndHash), nil
}

func (h finishedHash) hashForClientCertificate(signatureAndHash signatureAndHash, masterSecret []byte) ([]byte, crypto.Hash, error) {
	_ = "STUB: not implemented"
	return nil, *new(crypto.Hash), nil
}

func (h *finishedHash) discardHandshakeBuffer() { _ = "STUB: not implemented"; return }

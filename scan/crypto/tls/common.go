package tls

import (
	"container/list"
	"crypto"
	"crypto/x509"
	"io"
	"math/big"
	"sync"
	"time"
)

const (
	VersionSSL30 = 0x0300
	VersionTLS10 = 0x0301
	VersionTLS11 = 0x0302
	VersionTLS12 = 0x0303
)

const (
	maxPlaintext    = 16384
	maxCiphertext   = 16384 + 2048
	recordHeaderLen = 5
	maxHandshake    = 65536

	minVersion = VersionTLS10
	maxVersion = VersionTLS12
)

type recordType uint8

const (
	recordTypeChangeCipherSpec recordType = 20
	recordTypeAlert            recordType = 21
	recordTypeHandshake        recordType = 22
	recordTypeApplicationData  recordType = 23
)

const (
	typeClientHello        uint8 = 1
	typeServerHello        uint8 = 2
	typeNewSessionTicket   uint8 = 4
	typeCertificate        uint8 = 11
	typeServerKeyExchange  uint8 = 12
	typeCertificateRequest uint8 = 13
	typeServerHelloDone    uint8 = 14
	typeCertificateVerify  uint8 = 15
	typeClientKeyExchange  uint8 = 16
	typeFinished           uint8 = 20
	typeCertificateStatus  uint8 = 22
	typeNextProtocol       uint8 = 67
)

const (
	compressionNone uint8 = 0
)

const (
	extensionServerName          uint16 = 0
	extensionStatusRequest       uint16 = 5
	extensionSupportedCurves     uint16 = 10
	extensionSupportedPoints     uint16 = 11
	extensionSignatureAlgorithms uint16 = 13
	extensionALPN                uint16 = 16
	extensionSCT                 uint16 = 18
	extensionSessionTicket       uint16 = 35
	extensionNextProtoNeg        uint16 = 13172
	extensionRenegotiationInfo   uint16 = 0xff01
)

const (
	scsvRenegotiation uint16 = 0x00ff
)

type CurveID uint16

const (
	CurveP256 CurveID = 23
	CurveP384 CurveID = 24
	CurveP521 CurveID = 25
)

const (
	pointFormatUncompressed uint8 = 0
)

const (
	statusTypeOCSP uint8 = 1
)

const (
	certTypeRSASign    = 1
	certTypeDSSSign    = 2
	certTypeRSAFixedDH = 3
	certTypeDSSFixedDH = 4

	certTypeECDSASign      = 64
	certTypeRSAFixedECDH   = 65
	certTypeECDSAFixedECDH = 66
)

const (
	hashSHA1   uint8 = 2
	hashSHA256 uint8 = 4
	hashSHA384 uint8 = 5
)

const (
	signatureRSA   uint8 = 1
	signatureECDSA uint8 = 3
)

type signatureAndHash struct {
	hash, signature uint8
}

var supportedSignatureAlgorithms = []signatureAndHash{
	{hashSHA256, signatureRSA},
	{hashSHA256, signatureECDSA},
	{hashSHA384, signatureRSA},
	{hashSHA384, signatureECDSA},
	{hashSHA1, signatureRSA},
	{hashSHA1, signatureECDSA},
}

type ConnectionState struct {
	Version                     uint16
	HandshakeComplete           bool
	DidResume                   bool
	CipherSuite                 uint16
	NegotiatedProtocol          string
	NegotiatedProtocolIsMutual  bool
	ServerName                  string
	PeerCertificates            []*x509.Certificate
	VerifiedChains              [][]*x509.Certificate
	SignedCertificateTimestamps [][]byte
	OCSPResponse                []byte

	TLSUnique []byte
}

type ClientAuthType int

const (
	NoClientCert ClientAuthType = iota
	RequestClientCert
	RequireAnyClientCert
	VerifyClientCertIfGiven
	RequireAndVerifyClientCert
)

type ClientSessionState struct {
	sessionTicket      []uint8
	vers               uint16
	cipherSuite        uint16
	masterSecret       []byte
	serverCertificates []*x509.Certificate
	verifiedChains     [][]*x509.Certificate
}

type ClientSessionCache interface {
	Get(sessionKey string) (session *ClientSessionState, ok bool)

	Put(sessionKey string, cs *ClientSessionState)
}

type ClientHelloInfo struct {
	CipherSuites []uint16

	ServerName string

	SupportedCurves []CurveID

	SupportedPoints []uint8
}

type Config struct {
	Rand io.Reader

	Time func() time.Time

	Certificates []Certificate

	NameToCertificate map[string]*Certificate

	GetCertificate func(clientHello *ClientHelloInfo) (*Certificate, error)

	RootCAs *x509.CertPool

	NextProtos []string

	ServerName string

	ClientAuth ClientAuthType

	ClientCAs *x509.CertPool

	InsecureSkipVerify bool

	CipherSuites []uint16

	PreferServerCipherSuites bool

	SessionTicketsDisabled bool

	SessionTicketKey [32]byte

	ClientSessionCache ClientSessionCache

	MinVersion uint16

	MaxVersion uint16

	CurvePreferences []CurveID

	serverInitOnce sync.Once

	mutex sync.RWMutex

	sessionTicketKeys []ticketKey
}

const ticketKeyNameLen = 16

type ticketKey struct {
	keyName [ticketKeyNameLen]byte
	aesKey  [16]byte
	hmacKey [16]byte
}

func ticketKeyFromBytes(b [32]byte) (key ticketKey) {
	_ = "STUB: not implemented"
	return *new(ticketKey)
}

func (c *Config) serverInit() { _ = "STUB: not implemented"; return }

func (c *Config) ticketKeys() []ticketKey { _ = "STUB: not implemented"; return nil }

func (c *Config) SetSessionTicketKeys(keys [][32]byte) { _ = "STUB: not implemented"; return }

func (c *Config) rand() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (c *Config) time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *Config) cipherSuites() []uint16 { _ = "STUB: not implemented"; return nil }

func (c *Config) minVersion() uint16 { _ = "STUB: not implemented"; return 0 }

func (c *Config) maxVersion() uint16 { _ = "STUB: not implemented"; return 0 }

var defaultCurvePreferences = []CurveID{CurveP256, CurveP384, CurveP521}

func (c *Config) curvePreferences() []CurveID { _ = "STUB: not implemented"; return nil }

func (c *Config) mutualVersion(vers uint16) (uint16, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (c *Config) getCertificate(clientHello *ClientHelloInfo) (*Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Config) BuildNameToCertificate() { _ = "STUB: not implemented"; return }

type Certificate struct {
	Certificate [][]byte

	PrivateKey crypto.PrivateKey

	OCSPStaple []byte

	SignedCertificateTimestamps [][]byte

	Leaf *x509.Certificate
}

type record struct {
	contentType  recordType
	major, minor uint8
	payload      []byte
}

type handshakeMessage interface {
	marshal() []byte
	unmarshal([]byte) bool
}

type lruSessionCache struct {
	sync.Mutex

	m        map[string]*list.Element
	q        *list.List
	capacity int
}

type lruSessionCacheEntry struct {
	sessionKey string
	state      *ClientSessionState
}

func NewLRUClientSessionCache(capacity int) ClientSessionCache {
	_ = "STUB: not implemented"
	return *new(ClientSessionCache)
}

func (c *lruSessionCache) Put(sessionKey string, cs *ClientSessionState) {
	_ = "STUB: not implemented"
	return
}

func (c *lruSessionCache) Get(sessionKey string) (*ClientSessionState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type dsaSignature struct {
	R, S *big.Int
}

type ecdsaSignature dsaSignature

var emptyConfig Config

func defaultConfig() *Config { _ = "STUB: not implemented"; return nil }

var (
	once                   sync.Once
	varDefaultCipherSuites []uint16
)

func defaultCipherSuites() []uint16 { _ = "STUB: not implemented"; return nil }

func initDefaultCipherSuites() { _ = "STUB: not implemented"; return }

func unexpectedMessageError(wanted, got interface{}) error { _ = "STUB: not implemented"; return nil }

func isSupportedSignatureAndHash(sigHash signatureAndHash, sigHashes []signatureAndHash) bool {
	_ = "STUB: not implemented"
	return false
}

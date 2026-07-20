package tls

import (
	"bytes"
	"crypto/cipher"
	"crypto/x509"
	"errors"
	"io"
	"net"
	"sync"
	"time"
)

type Conn struct {
	conn     net.Conn
	isClient bool

	handshakeMutex    sync.Mutex
	handshakeErr      error
	vers              uint16
	haveVers          bool
	config            *Config
	handshakeComplete bool
	didResume         bool
	cipherSuite       uint16
	ocspResponse      []byte
	scts              [][]byte
	peerCertificates  []*x509.Certificate

	verifiedChains [][]*x509.Certificate

	serverName string

	firstFinished [12]byte

	clientProtocol         string
	clientProtocolFallback bool

	in, out  halfConn
	rawInput *block
	input    *block
	hand     bytes.Buffer

	activeCall int32

	tmp [16]byte
}

func (c *Conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

type halfConn struct {
	sync.Mutex

	err            error
	version        uint16
	cipher         interface{}
	mac            macFunction
	seq            [8]byte
	bfree          *block
	additionalData [13]byte

	nextCipher interface{}
	nextMac    macFunction

	inDigestBuf, outDigestBuf []byte
}

func (hc *halfConn) setErrorLocked(err error) error { _ = "STUB: not implemented"; return nil }

func (hc *halfConn) error() error { _ = "STUB: not implemented"; return nil }

func (hc *halfConn) prepareCipherSpec(version uint16, cipher interface{}, mac macFunction) {
	_ = "STUB: not implemented"
	return
}

func (hc *halfConn) changeCipherSpec() error { _ = "STUB: not implemented"; return nil }

func (hc *halfConn) incSeq() { _ = "STUB: not implemented"; return }

func (hc *halfConn) resetSeq() { _ = "STUB: not implemented"; return }

func removePadding(payload []byte) ([]byte, byte) { _ = "STUB: not implemented"; return nil, 0 }

func removePaddingSSL30(payload []byte) ([]byte, byte) { _ = "STUB: not implemented"; return nil, 0 }

func roundUp(a, b int) int { _ = "STUB: not implemented"; return 0 }

type cbcMode interface {
	cipher.BlockMode
	SetIV([]byte)
}

func (hc *halfConn) decrypt(b *block) (ok bool, prefixLen int, alertValue alert) {
	_ = "STUB: not implemented"
	return false, 0, *new(alert)
}

func padToBlockSize(payload []byte, blockSize int) (prefix, finalBlock []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hc *halfConn) encrypt(b *block, explicitIVLen int) (bool, alert) {
	_ = "STUB: not implemented"
	return false, *new(alert)
}

type block struct {
	data []byte
	off  int
	link *block
}

func (b *block) resize(n int) { _ = "STUB: not implemented"; return }

func (b *block) reserve(n int) { _ = "STUB: not implemented"; return }

func (b *block) readFromUntil(r io.Reader, n int) error { _ = "STUB: not implemented"; return nil }

func (b *block) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (hc *halfConn) newBlock() *block { _ = "STUB: not implemented"; return nil }

func (hc *halfConn) freeBlock(b *block) { _ = "STUB: not implemented"; return }

func (hc *halfConn) splitBlock(b *block, n int) (*block, *block) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RecordHeaderError struct {
	Msg string

	RecordHeader [5]byte
}

func (e RecordHeaderError) Error() string { _ = "STUB: not implemented"; return "" }

func (c *Conn) newRecordHeaderError(msg string) (err RecordHeaderError) {
	_ = "STUB: not implemented"
	return *new(RecordHeaderError)
}

func (c *Conn) readRecord(want recordType) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) sendAlertLocked(err alert) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) sendAlert(err alert) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) writeRecord(typ recordType, data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Conn) readHandshake() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

var errClosed = errors.New("crypto/tls: use of closed connection")

func (c *Conn) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) Handshake() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) ConnectionState() ConnectionState {
	_ = "STUB: not implemented"
	return *new(ConnectionState)
}

func (c *Conn) OCSPResponse() []byte { _ = "STUB: not implemented"; return nil }

func (c *Conn) VerifyHostname(host string) error { _ = "STUB: not implemented"; return nil }

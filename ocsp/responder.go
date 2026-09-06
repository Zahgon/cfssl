package ocsp

import (
	"crypto"
	"errors"
	"net/http"
	"time"

	"github.com/cloudflare/cfssl/certdb"
	"github.com/jmhodges/clock"
	"golang.org/x/crypto/ocsp"
)

var (
	malformedRequestErrorResponse = []byte{0x30, 0x03, 0x0A, 0x01, 0x01}
	internalErrorErrorResponse    = []byte{0x30, 0x03, 0x0A, 0x01, 0x02}
	tryLaterErrorResponse         = []byte{0x30, 0x03, 0x0A, 0x01, 0x03}
	sigRequredErrorResponse       = []byte{0x30, 0x03, 0x0A, 0x01, 0x05}
	unauthorizedErrorResponse     = []byte{0x30, 0x03, 0x0A, 0x01, 0x06}

	ErrNotFound = errors.New("Request OCSP Response not found")
)

type Source interface {
	Response(*ocsp.Request) ([]byte, http.Header, error)
}

type InMemorySource map[string][]byte

func (src InMemorySource) Response(request *ocsp.Request) ([]byte, http.Header, error) {
	_ = "STUB: not implemented"
	return nil, *new(http.Header), nil
}

type DBSource struct {
	Accessor certdb.Accessor
}

func NewDBSource(dbAccessor certdb.Accessor) Source { _ = "STUB: not implemented"; return *new(Source) }

func (src DBSource) Response(req *ocsp.Request) ([]byte, http.Header, error) {
	_ = "STUB: not implemented"
	return nil, *new(http.Header), nil
}

func NewSourceFromFile(responseFile string) (Source, error) {
	_ = "STUB: not implemented"
	return *new(Source), nil
}

func NewSourceFromDB(DBConfigFile string) (Source, error) {
	_ = "STUB: not implemented"
	return *new(Source), nil
}

type Stats interface {
	ResponseStatus(ocsp.ResponseStatus)
}

type Responder struct {
	Source Source
	stats  Stats
	clk    clock.Clock
}

func NewResponder(source Source, stats Stats) *Responder { _ = "STUB: not implemented"; return nil }

func overrideHeaders(response http.ResponseWriter, headers http.Header) {
	_ = "STUB: not implemented"
	return
}

type logEvent struct {
	IP       string        `json:"ip,omitempty"`
	UA       string        `json:"ua,omitempty"`
	Method   string        `json:"method,omitempty"`
	Path     string        `json:"path,omitempty"`
	Body     string        `json:"body,omitempty"`
	Received time.Time     `json:"received,omitempty"`
	Took     time.Duration `json:"took,omitempty"`
	Headers  http.Header   `json:"headers,omitempty"`

	Serial         string `json:"serial,omitempty"`
	IssuerKeyHash  string `json:"issuerKeyHash,omitempty"`
	IssuerNameHash string `json:"issuerNameHash,omitempty"`
	HashAlg        string `json:"hashAlg,omitempty"`
}

var hashToString = map[crypto.Hash]string{
	crypto.SHA1:   "SHA1",
	crypto.SHA256: "SHA256",
	crypto.SHA384: "SHA384",
	crypto.SHA512: "SHA512",
}

func (rs Responder) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	_ = "STUB: not implemented"
	return
}

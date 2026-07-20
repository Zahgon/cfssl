package scan

import (
	"crypto/x509"
	"net"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/cloudflare/cfssl/scan/crypto/tls"
)

var (
	Network = "tcp"

	Dialer = &net.Dialer{Timeout: time.Second}

	Client = &http.Client{Transport: &http.Transport{Dial: Dialer.Dial}}

	RootCAs *x509.CertPool
)

type Grade int

const (
	Bad Grade = iota

	Warning

	Good

	Skipped
)

func (g Grade) String() string { _ = "STUB: not implemented"; return "" }

type Output interface{}

func multiscan(host string, scan func(string) (Grade, Output, error)) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

type Scanner struct {
	Description string `json:"description"`

	scan func(string, string) (Grade, Output, error)
}

func (s *Scanner) Scan(addr, hostname string) (Grade, Output, error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

type Family struct {
	Description string `json:"description"`

	Scanners map[string]*Scanner `json:"scanners"`
}

type FamilySet map[string]*Family

var Default = FamilySet{
	"Connectivity": Connectivity,
	"TLSHandshake": TLSHandshake,
	"TLSSession":   TLSSession,
	"PKI":          PKI,
	"Broad":        Broad,
}

type ScannerResult struct {
	Grade  string `json:"grade"`
	Output Output `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

type FamilyResult map[string]ScannerResult

type Result struct {
	Family, Scanner string
	ScannerResult
}

type context struct {
	sync.WaitGroup
	addr, hostname              string
	familyRegexp, scannerRegexp *regexp.Regexp
	resultChan                  chan *Result
}

func newContext(addr, hostname string, familyRegexp, scannerRegexp *regexp.Regexp, numFamilies int) *context {
	_ = "STUB: not implemented"
	return nil
}

type familyContext struct {
	sync.WaitGroup
	ctx *context
}

func (ctx *context) newfamilyContext(numScanners int) *familyContext {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *context) copyResults(timeout time.Duration) map[string]FamilyResult {
	_ = "STUB: not implemented"
	return nil
}

func (familyCtx *familyContext) runScanner(familyName, scannerName string, scanner *Scanner) {
	_ = "STUB: not implemented"
	return
}

func (fs FamilySet) RunScans(host, ip, family, scanner string, timeout time.Duration) (map[string]FamilyResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadRootCAs(caBundleFile string) (err error) { _ = "STUB: not implemented"; return nil }

func defaultTLSConfig(hostname string) *tls.Config { _ = "STUB: not implemented"; return nil }

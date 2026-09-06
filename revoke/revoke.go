package revoke

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"net/http"
	"sync"

	"golang.org/x/crypto/ocsp"
)

var HTTPClient = http.DefaultClient

var HardFail = false

var CRLSet = map[string]*pkix.CertificateList{}
var crlLock = new(sync.Mutex)

func ldapURL(url string) bool { _ = "STUB: not implemented"; return false }

func revCheck(cert *x509.Certificate) (revoked, ok bool, err error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func fetchCRL(url string) (*pkix.CertificateList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getIssuer(cert *x509.Certificate) *x509.Certificate { _ = "STUB: not implemented"; return nil }

func certIsRevokedCRL(cert *x509.Certificate, url string) (revoked, ok bool, err error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func VerifyCertificate(cert *x509.Certificate) (revoked, ok bool) {
	_ = "STUB: not implemented"
	return false, false
}

func VerifyCertificateError(cert *x509.Certificate) (revoked, ok bool, err error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func fetchRemote(url string) (*x509.Certificate, error) { _ = "STUB: not implemented"; return nil, nil }

var ocspOpts = ocsp.RequestOptions{
	Hash: crypto.SHA1,
}

func certIsRevokedOCSP(leaf *x509.Certificate, strict bool) (revoked, ok bool, e error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func sendOCSPRequest(server string, req []byte, leaf, issuer *x509.Certificate) (*ocsp.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var crlRead = io.ReadAll

func SetCRLFetcher(fn func(io.Reader) ([]byte, error)) { _ = "STUB: not implemented"; return }

var remoteRead = io.ReadAll

func SetRemoteFetcher(fn func(io.Reader) ([]byte, error)) { _ = "STUB: not implemented"; return }

var ocspRead = io.ReadAll

func SetOCSPFetcher(fn func(io.Reader) ([]byte, error)) { _ = "STUB: not implemented"; return }

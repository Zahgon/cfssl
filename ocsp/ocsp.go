package ocsp

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"time"

	"golang.org/x/crypto/ocsp"
)

var revocationReasonCodes = map[string]int{
	"unspecified":          ocsp.Unspecified,
	"keycompromise":        ocsp.KeyCompromise,
	"cacompromise":         ocsp.CACompromise,
	"affiliationchanged":   ocsp.AffiliationChanged,
	"superseded":           ocsp.Superseded,
	"cessationofoperation": ocsp.CessationOfOperation,
	"certificatehold":      ocsp.CertificateHold,
	"removefromcrl":        ocsp.RemoveFromCRL,
	"privilegewithdrawn":   ocsp.PrivilegeWithdrawn,
	"aacompromise":         ocsp.AACompromise,
}

var StatusCode = map[string]int{
	"good":    ocsp.Good,
	"revoked": ocsp.Revoked,
	"unknown": ocsp.Unknown,
}

type SignRequest struct {
	Certificate *x509.Certificate
	Status      string
	Reason      int
	RevokedAt   time.Time
	Extensions  []pkix.Extension

	IssuerHash crypto.Hash

	ThisUpdate *time.Time

	NextUpdate *time.Time
}

type Signer interface {
	Sign(req SignRequest) ([]byte, error)
}

type StandardSigner struct {
	issuer    *x509.Certificate
	responder *x509.Certificate
	key       crypto.Signer
	interval  time.Duration
}

func ReasonStringToCode(reason string) (reasonCode int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func NewSignerFromFile(issuerFile, responderFile, keyFile string, interval time.Duration) (Signer, error) {
	_ = "STUB: not implemented"
	return *new(Signer), nil
}

func NewSigner(issuer, responder *x509.Certificate, key crypto.Signer, interval time.Duration) (Signer, error) {
	_ = "STUB: not implemented"
	return *new(Signer), nil
}

func (s StandardSigner) Sign(req SignRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

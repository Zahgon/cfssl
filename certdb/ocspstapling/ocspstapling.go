package ocspstapling

import (
	"crypto"
	"crypto/x509"
	"encoding/asn1"

	"github.com/cloudflare/cfssl/certdb"
	ct "github.com/google/certificate-transparency-go"
)

var sctExtOid = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 4, 5}

func StapleSCTList(acc certdb.Accessor, serial, aki string, scts []ct.SignedCertificateTimestamp,
	responderCert, issuer *x509.Certificate, priv crypto.Signer) error {
	_ = "STUB: not implemented"
	return nil
}

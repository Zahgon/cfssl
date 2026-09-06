package pkcs7

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
)

type signedData struct {
	Version          int
	DigestAlgorithms asn1.RawValue
	ContentInfo      asn1.RawValue
	Certificates     asn1.RawValue `asn1:"optional" asn1:"tag:0"`
	Crls             asn1.RawValue `asn1:"optional"`
	SignerInfos      asn1.RawValue
}

type initPKCS7 struct {
	Raw         asn1.RawContent
	ContentType asn1.ObjectIdentifier
	Content     asn1.RawValue `asn1:"tag:0,explicit,optional"`
}

const (
	ObjIDData          = "1.2.840.113549.1.7.1"
	ObjIDSignedData    = "1.2.840.113549.1.7.2"
	ObjIDEncryptedData = "1.2.840.113549.1.7.6"
)

type PKCS7 struct {
	Raw         asn1.RawContent
	ContentInfo string
	Content     Content
}

type Content struct {
	Data          []byte
	SignedData    SignedData
	EncryptedData EncryptedData
}

type SignedData struct {
	Raw          asn1.RawContent
	Version      int
	Certificates []*x509.Certificate
	Crl          *pkix.CertificateList
}

type Data struct {
	Bytes []byte
}

type EncryptedData struct {
	Raw                  asn1.RawContent
	Version              int
	EncryptedContentInfo EncryptedContentInfo
}

type EncryptedContentInfo struct {
	Raw                        asn1.RawContent
	ContentType                asn1.ObjectIdentifier
	ContentEncryptionAlgorithm pkix.AlgorithmIdentifier
	EncryptedContent           []byte `asn1:"tag:0,optional"`
}

func ParsePKCS7(raw []byte) (msg *PKCS7, err error) { _ = "STUB: not implemented"; return nil, nil }

package ca

type CertificateAuthority interface {
	SignCSR(csrPEM []byte) (cert []byte, err error)

	CACertificate() (cert []byte, err error)
}

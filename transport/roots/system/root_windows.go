package system

import (
	"crypto/x509"
	"syscall"
)

func createStoreContext(leaf *Certificate, opts *VerifyOptions) (*syscall.CertContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractSimpleChain(simpleChain **syscall.CertSimpleChain, count int) (chain []*Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkChainTrustStatus(c *Certificate, chainCtx *syscall.CertChainContext) error {
	_ = "STUB: not implemented"
	return nil
}

func checkChainSSLServerPolicy(c *Certificate, chainCtx *syscall.CertChainContext, opts *VerifyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func initSystemRoots() []*x509.Certificate { _ = "STUB: not implemented"; return nil }

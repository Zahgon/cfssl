package bundler

import (
	"crypto"
	"crypto/x509"
	"net/http"
)

var IntermediateStash string

var HTTPClient = http.DefaultClient

type BundleFlavor string

const (
	Optimal BundleFlavor = "optimal"

	Ubiquitous BundleFlavor = "ubiquitous"

	Force BundleFlavor = "force"
)

const (
	sha2Warning          = "The bundle contains certificates signed with advanced hash functions such as SHA2, which are problematic for certain operating systems, e.g. Windows XP SP2."
	ecdsaWarning         = "The bundle contains ECDSA signatures, which are problematic for certain operating systems, e.g. Windows XP, Android 2.2 and Android 2.3."
	expiringWarningStub  = "The bundle is expiring within 30 days."
	untrustedWarningStub = "The bundle may not be trusted by the following platform(s):"
	ubiquityWarning      = "Unable to measure bundle ubiquity: No platform metadata present."
)

type Bundler struct {
	RootPool         *x509.CertPool
	IntermediatePool *x509.CertPool
	KnownIssuers     map[string]bool
	opts             options
}

type options struct {
	keyUsages []x509.ExtKeyUsage
}

var defaultOptions = options{
	keyUsages: []x509.ExtKeyUsage{
		x509.ExtKeyUsageAny,
	},
}

type Option func(*options)

func WithKeyUsages(usages ...x509.ExtKeyUsage) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func NewBundler(caBundleFile, intBundleFile string, opt ...Option) (*Bundler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBundlerFromPEM(caBundlePEM, intBundlePEM []byte, opt ...Option) (*Bundler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bundler) VerifyOptions() x509.VerifyOptions {
	_ = "STUB: not implemented"
	return *new(x509.VerifyOptions)
}

func (b *Bundler) BundleFromFile(bundleFile, keyFile string, flavor BundleFlavor, password string) (*Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bundler) BundleFromPEMorDER(certsRaw, keyPEM []byte, flavor BundleFlavor, password string) (*Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bundler) BundleFromRemote(serverName, ip string, flavor BundleFlavor) (*Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type fetchedIntermediate struct {
	Cert *x509.Certificate
	Name string
}

func fetchRemoteCertificate(certURL string) (fi *fetchedIntermediate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func reverse(certs []*x509.Certificate) []*x509.Certificate { _ = "STUB: not implemented"; return nil }

func partialVerify(certs []*x509.Certificate) bool { _ = "STUB: not implemented"; return false }

func isSelfSigned(cert *x509.Certificate) bool { _ = "STUB: not implemented"; return false }

func isChainRootNode(cert *x509.Certificate) bool { _ = "STUB: not implemented"; return false }

func (b *Bundler) verifyChain(chain []*fetchedIntermediate) bool {
	_ = "STUB: not implemented"
	return false
}

func constructCertFileName(cert *x509.Certificate) string { _ = "STUB: not implemented"; return "" }

func (b *Bundler) fetchIntermediates(certs []*x509.Certificate) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bundler) Bundle(certs []*x509.Certificate, key crypto.Signer, flavor BundleFlavor) (*Bundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkExpiringCerts(chain []*x509.Certificate) (expiringIntermediates []int) {
	_ = "STUB: not implemented"
	return nil
}

func getSKIs(chain []*x509.Certificate, indices []int) (skis []string) {
	_ = "STUB: not implemented"
	return nil
}

func expirationWarning(expiringIntermediates []int) (ret string) {
	_ = "STUB: not implemented"
	return ""
}

func untrustedPlatformsWarning(platforms []string) string { _ = "STUB: not implemented"; return "" }

func optimalChains(chains [][]*x509.Certificate) [][]*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func ubiquitousChains(chains [][]*x509.Certificate) [][]*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func diff(chain1, chain2 []*x509.Certificate) bool { _ = "STUB: not implemented"; return false }

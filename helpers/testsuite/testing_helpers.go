package testsuite

import (
	"os"
	"testing"

	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/csr"
)

type CFSSLServerData struct {
	CA        []byte
	CABundle  []byte
	CAKey     []byte
	IntBundle []byte
}

type CFSSLServer struct {
	process   *os.Process
	tempFiles []string
}

func StartCFSSLServer(address string, portNumber int, serverData CFSSLServerData) (*CFSSLServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (server *CFSSLServer) Kill() error { _ = "STUB: not implemented"; return nil }

func CreateCertificateChain(requests []csr.CertificateRequest) (certChain []byte, key []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateSelfSignedCert(request csr.CertificateRequest) (encodedCert, encodedKey []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func SignCertificate(request csr.CertificateRequest, signerCert, signerKey []byte) (encodedCert, encodedKey []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func createTempFile(data []byte) (fileName string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func checkCLIOutput(CLIOutput []byte) error { _ = "STUB: not implemented"; return nil }

func cleanCLIOutput(CLIOutput []byte, item string) (cleanedOutput []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewConfig(t *testing.T, configBytes []byte) *config.Config {
	_ = "STUB: not implemented"
	return nil
}

type CSRTest struct {
	File    string
	KeyAlgo string
	KeyLen  int

	ErrorCallback func(*testing.T, error)
}

var CSRTests = []CSRTest{
	{
		File:          "../../signer/local/testdata/rsa2048.csr",
		KeyAlgo:       "rsa",
		KeyLen:        2048,
		ErrorCallback: nil,
	},
	{
		File:          "../../signer/local/testdata/rsa3072.csr",
		KeyAlgo:       "rsa",
		KeyLen:        3072,
		ErrorCallback: nil,
	},
	{
		File:          "../../signer/local/testdata/rsa4096.csr",
		KeyAlgo:       "rsa",
		KeyLen:        4096,
		ErrorCallback: nil,
	},
	{
		File:          "../../signer/local/testdata/ecdsa256.csr",
		KeyAlgo:       "ecdsa",
		KeyLen:        256,
		ErrorCallback: nil,
	},
	{
		File:          "../../signer/local/testdata/ecdsa384.csr",
		KeyAlgo:       "ecdsa",
		KeyLen:        384,
		ErrorCallback: nil,
	},
	{
		File:          "../../signer/local/testdata/ecdsa521.csr",
		KeyAlgo:       "ecdsa",
		KeyLen:        521,
		ErrorCallback: nil,
	},
}

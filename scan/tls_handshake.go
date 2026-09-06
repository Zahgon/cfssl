package scan

import (
	"errors"

	"github.com/cloudflare/cfssl/scan/crypto/tls"
)

var errHelloFailed = errors.New("Handshake failed in sayHello")

var TLSHandshake = &Family{
	Description: "Scans for host's SSL/TLS version and cipher suite negotiation",
	Scanners: map[string]*Scanner{
		"CipherSuite": {
			"Determines host's cipher suites accepted and preferred order",
			cipherSuiteScan,
		},
		"SigAlgs": {
			"Determines host's accepted signature and hash algorithms",
			sigAlgsScan,
		},
		"CertsBySigAlgs": {
			"Determines host's certificate signature algorithm matching client's accepted signature and hash algorithms",
			certSigAlgsScan,
		},
		"CertsByCiphers": {
			"Determines host's certificate signature algorithm matching client's accepted ciphers",
			certSigAlgsScanByCipher,
		},
		"ECCurves": {
			"Determines the host's ec curve support for TLS 1.2",
			ecCurveScan,
		},
	},
}

func getCipherIndex(ciphers []uint16, serverCipher uint16) (cipherIndex int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getCurveIndex(curves []tls.CurveID, serverCurve tls.CurveID) (curveIndex int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func sayHello(addr, hostname string, ciphers []uint16, curves []tls.CurveID, vers uint16, sigAlgs []tls.SignatureAndHash) (cipherIndex, curveIndex int, certs [][]byte, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil, nil
}

func allCiphersIDs() []uint16 { _ = "STUB: not implemented"; return nil }

func allECDHECiphersIDs() []uint16 { _ = "STUB: not implemented"; return nil }

func allCurvesIDs() []tls.CurveID { _ = "STUB: not implemented"; return nil }

type cipherDatum struct {
	versionID uint16
	curves    []tls.CurveID
}

type cipherVersions struct {
	cipherID uint16
	data     []cipherDatum
}

type cipherVersionList []cipherVersions

func (cvList cipherVersionList) String() string { _ = "STUB: not implemented"; return "" }

func (cvList cipherVersionList) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func doCurveScan(addr, hostname string, vers, cipherID uint16, ciphers []uint16) (supportedCurves []tls.CurveID, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cipherSuiteScan(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

func sigAlgsScan(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

func certSigAlgsScan(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

func certSigAlgsScanByCipher(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

func ecCurveScan(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

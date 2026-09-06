package scan

import (
	"net"
)

var Connectivity = &Family{
	Description: "Scans for basic connectivity with the host through DNS and TCP/TLS dials",
	Scanners: map[string]*Scanner{
		"DNSLookup": {
			"Host can be resolved through DNS",
			dnsLookupScan,
		},
		"CloudFlareStatus": {
			"Host is on CloudFlare",
			onCloudFlareScan,
		},
		"TCPDial": {
			"Host accepts TCP connection",
			tcpDialScan,
		},
		"TLSDial": {
			"Host can perform TLS handshake",
			tlsDialScan,
		},
	},
}

func dnsLookupScan(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

var (
	cfNets    []*net.IPNet
	cfNetsErr error
)

func initOnCloudFlareScan() ([]*net.IPNet, error) { _ = "STUB: not implemented"; return nil, nil }

func onCloudFlareScan(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

func tcpDialScan(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

func tlsDialScan(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

package scan

import (
	"time"
)

var Broad = &Family{
	Description: "Large scale scans of TLS hosts",
	Scanners: map[string]*Scanner{
		"IntermediateCAs": {
			"Scans a CIDR IP range for unknown Intermediate CAs",
			intermediateCAScan,
		},
	},
}

func incrementBytes(bytes []byte) { _ = "STUB: not implemented"; return }

var (
	caBundleFile  = "/etc/cfssl/ca-bundle.crt"
	intBundleFile = "/etc/cfssl/int-bundle.crt"
	numWorkers    = 32
	timeout       = time.Second
)

func intermediateCAScan(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

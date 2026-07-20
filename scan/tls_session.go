package scan

var TLSSession = &Family{
	Description: "Scans host's implementation of TLS session resumption using session tickets/session IDs",
	Scanners: map[string]*Scanner{
		"SessionResume": {
			"Host is able to resume sessions across all addresses",
			sessionResumeScan,
		},
	},
}

func sessionResumeScan(addr, hostname string) (grade Grade, output Output, err error) {
	_ = "STUB: not implemented"
	return *new(Grade), *new(Output), nil
}

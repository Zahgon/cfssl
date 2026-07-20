package core

import (
	"time"

	"github.com/cloudflare/cfssl/csr"
)

type Root struct {
	Type string `json:"type"`

	Metadata map[string]string `json:"metadata"`
}

type Identity struct {
	Request *csr.CertificateRequest `json:"request"`

	Roots []*Root `json:"roots"`

	ClientRoots []*Root `json:"client_roots"`

	Profiles map[string]map[string]string `json:"profiles"`
}

var DefaultBefore = 24 * time.Hour

var CipherSuites = []uint16{

	0xc030,
	0xc02c,
	0xc02f,
	0xc02b,
}

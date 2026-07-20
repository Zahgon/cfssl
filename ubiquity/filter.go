package ubiquity

import (
	"crypto/x509"
)

type RankingFunc func(chain1, chain2 []*x509.Certificate) int

func Filter(chains [][]*x509.Certificate, f RankingFunc) [][]*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

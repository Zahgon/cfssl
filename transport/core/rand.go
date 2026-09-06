package core

import (
	"github.com/cloudflare/cfssl/log"
)

var seeded bool

func seed() error { _ = "STUB: not implemented"; return nil }

func init() {
	err := seed()
	if err != nil {
		log.Errorf("seeding mrand failed: %v", err)
	}
}

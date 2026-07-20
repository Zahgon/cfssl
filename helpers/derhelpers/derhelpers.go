package derhelpers

import (
	"crypto"
)

func ParsePrivateKeyDER(keyDER []byte) (key crypto.Signer, err error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

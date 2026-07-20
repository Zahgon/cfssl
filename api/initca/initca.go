package initca

import (
	"net/http"
)

type NewCA struct {
	Key  string `json:"private_key"`
	Cert string `json:"certificate"`
}

func initialCAHandler(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

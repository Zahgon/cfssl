package health

import (
	"net/http"
)

type Response struct {
	Healthy bool `json:"healthy"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHealthCheck() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

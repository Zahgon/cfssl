package scan

import (
	"net/http"
)

func scanHandler(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHandler(caBundleFile string) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func scanInfoHandler(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInfoHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

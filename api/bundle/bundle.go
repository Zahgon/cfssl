package bundle

import (
	"net/http"

	"github.com/cloudflare/cfssl/bundler"
)

type Handler struct {
	bundler *bundler.Bundler
}

func NewHandler(caBundleFile, intBundleFile string) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

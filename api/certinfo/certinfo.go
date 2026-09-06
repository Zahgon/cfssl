package certinfo

import (
	"net/http"

	"github.com/cloudflare/cfssl/certdb"
)

type Handler struct {
	dbAccessor certdb.Accessor
}

func NewHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func NewAccessorHandler(dbAccessor certdb.Accessor) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) (err error) {
	_ = "STUB: not implemented"
	return nil
}

package whitelist

import (
	"net"
	"net/http"
)

func NetConnLookup(conn net.Conn) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func HTTPRequestLookup(req *http.Request) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

type Handler struct {
	allowHandler http.Handler
	denyHandler  http.Handler
	whitelist    ACL
}

func NewHandler(allow, deny http.Handler, acl ACL) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

type HandlerFunc struct {
	allow     func(http.ResponseWriter, *http.Request)
	deny      func(http.ResponseWriter, *http.Request)
	whitelist ACL
}

func NewHandlerFunc(allow, deny func(http.ResponseWriter, *http.Request), acl ACL) (*HandlerFunc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

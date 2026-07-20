package gencrl

import (
	"net/http"
)

type jsonCRLRequest struct {
	Certificate  string   `json:"certificate"`
	SerialNumber []string `json:"serialNumber"`
	PrivateKey   string   `json:"issuingKey"`
	ExpiryTime   string   `json:"expireTime"`
}

func gencrlHandler(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

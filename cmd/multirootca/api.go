package main

import (
	"net/http"

	"github.com/cloudflare/cfssl/signer"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type SignatureResponse struct {
	Certificate string `json:"certificate"`
}

type filter func(string, *signer.SignRequest) bool

var filters = map[string][]filter{}
var (
	requests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_total",
			Help: "How many requests for each operation type and signer were succesfully processed.",
		},
		[]string{"operation", "signer"},
	)
	erroredRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_errored_total",
			Help: "How many requests for each operation type resulted in an error.",
		},
		[]string{"operation", "signer"},
	)
	badInputs = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "bad_inputs_total",
			Help: "How many times the input was malformed or not allowed.",
		},
		[]string{"operation"},
	)
)

const (
	signOperation = "sign"
)

func fail(w http.ResponseWriter, req *http.Request, status, code int, msg, ad string) {
	_ = "STUB: not implemented"
	return
}

func dispatchRequest(w http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func metricsDisallowed(w http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

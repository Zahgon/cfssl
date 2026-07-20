package main

import (
	"flag"
	"net"
	"net/http"

	"github.com/cloudflare/cfssl/api/info"
	"github.com/cloudflare/cfssl/log"
	"github.com/cloudflare/cfssl/multiroot/config"
	"github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/whitelist"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func parseSigner(root *config.Root) (signer.Signer, error) {
	_ = "STUB: not implemented"
	return *new(signer.Signer), nil
}

var (
	defaultLabel string
	signers      = map[string]signer.Signer{}
	whitelists   = map[string]whitelist.NetACL{}
)

func main() {
	flagAddr := flag.String("a", ":8888", "listening address")
	flagRootFile := flag.String("roots", "", "configuration file specifying root keys")
	flagDefaultLabel := flag.String("l", "", "specify a default label")
	flagEndpointCert := flag.String("tls-cert", "", "server certificate")
	flagEndpointKey := flag.String("tls-key", "", "server private key")
	flag.IntVar(&log.Level, "loglevel", log.LevelInfo, "Log level (0 = DEBUG, 5 = FATAL)")
	flag.Parse()

	if *flagRootFile == "" {
		log.Fatal("no root file specified")
	}

	roots, err := config.Parse(*flagRootFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for label, root := range roots {
		s, err := parseSigner(root)
		if err != nil {
			log.Criticalf("%v", err)
		}
		signers[label] = s
		if root.ACL != nil {
			whitelists[label] = root.ACL
		}
		log.Info("loaded signer ", label)
	}

	defaultLabel = *flagDefaultLabel

	infoHandler, err := info.NewMultiHandler(signers, defaultLabel)
	if err != nil {
		log.Criticalf("%v", err)
	}

	var localhost = whitelist.NewBasic()
	localhost.Add(net.ParseIP("127.0.0.1"))
	localhost.Add(net.ParseIP("::1"))

	http.HandleFunc("/api/v1/cfssl/authsign", dispatchRequest)
	http.Handle("/api/v1/cfssl/info", infoHandler)
	http.Handle("/metrics", promhttp.Handler())

	if *flagEndpointCert == "" && *flagEndpointKey == "" {
		log.Info("Now listening on ", *flagAddr)
		log.Fatal(http.ListenAndServe(*flagAddr, nil))
	} else {

		log.Info("Now listening on https:// ", *flagAddr)
		log.Fatal(http.ListenAndServeTLS(*flagAddr, *flagEndpointCert, *flagEndpointKey, nil))
	}

}

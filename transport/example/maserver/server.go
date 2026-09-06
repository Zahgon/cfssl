package main

import (
	"encoding/json"
	"flag"
	"net"
	"os"

	"github.com/cloudflare/cfssl/log"
	"github.com/cloudflare/cfssl/transport"
	"github.com/cloudflare/cfssl/transport/core"
	"github.com/cloudflare/cfssl/transport/example/exlib"
)

func main() {
	var addr, conf string
	flag.StringVar(&addr, "a", "127.0.0.1:9876", "`address` of server")
	flag.StringVar(&conf, "f", "server.json", "config `file` to use")
	flag.Parse()

	var id = new(core.Identity)
	data, err := os.ReadFile(conf)
	if err != nil {
		exlib.Err(1, err, "reading config file")
	}

	err = json.Unmarshal(data, id)
	if err != nil {
		exlib.Err(1, err, "parsing config file")
	}

	tr, err := transport.New(exlib.Before, id)
	if err != nil {
		exlib.Err(1, err, "creating transport")
	}

	l, err := transport.Listen(addr, tr)
	if err != nil {
		exlib.Err(1, err, "setting up listener")
	}

	var errChan = make(chan error, 0)
	go func(ec <-chan error) {
		for {
			err, ok := <-ec
			if !ok {
				log.Warning("error channel closed, future errors will not be reported")
				break
			}
			log.Errorf("auto update error: %v", err)
		}
	}(errChan)

	log.Info("setting up auto-update")
	go l.AutoUpdate(nil, errChan)

	log.Info("listening on ", addr)
	exlib.Warn(serve(l), "serving listener")
}

func connHandler(conn net.Conn) { _ = "STUB: not implemented"; return }

func serve(l net.Listener) error { _ = "STUB: not implemented"; return nil }

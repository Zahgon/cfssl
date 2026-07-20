package main

import (
	"crypto/x509"
	"flag"
	"sync"
)

func worker(paths chan string, bundler chan *x509.Certificate, pool *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

func supervisor(paths chan string, bundler chan *x509.Certificate, numWorkers int) {
	_ = "STUB: not implemented"
	return
}

func makeBundle(filename string, bundler chan *x509.Certificate) { _ = "STUB: not implemented"; return }

func scanFiles(paths chan string) { _ = "STUB: not implemented"; return }

func main() {
	bundleFile := flag.String("f", "cert-bundle.crt", "path to store certificate bundle")
	numWorkers := flag.Int("nw", 4, "number of workers")
	flag.Parse()

	paths := make(chan string)
	bundler := make(chan *x509.Certificate)

	go supervisor(paths, bundler, *numWorkers)
	go scanFiles(paths)

	makeBundle(*bundleFile, bundler)
}

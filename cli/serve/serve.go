package serve

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"

	"github.com/cloudflare/cfssl/api"
	"github.com/cloudflare/cfssl/api/bundle"
	"github.com/cloudflare/cfssl/api/certadd"
	"github.com/cloudflare/cfssl/api/certinfo"
	"github.com/cloudflare/cfssl/api/crl"
	"github.com/cloudflare/cfssl/api/gencrl"
	"github.com/cloudflare/cfssl/api/generator"
	"github.com/cloudflare/cfssl/api/health"
	"github.com/cloudflare/cfssl/api/info"
	"github.com/cloudflare/cfssl/api/initca"
	apiocsp "github.com/cloudflare/cfssl/api/ocsp"
	"github.com/cloudflare/cfssl/api/revoke"
	"github.com/cloudflare/cfssl/api/scan"
	"github.com/cloudflare/cfssl/api/signhandler"
	certsql "github.com/cloudflare/cfssl/certdb/sql"
	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/ocsp"
	"github.com/cloudflare/cfssl/signer"

	"github.com/jmoiron/sqlx"
)

var serverUsageText = `cfssl serve -- set up a HTTP server handles CF SSL requests

Usage of serve:
        cfssl serve [-address address] [-min-tls-version version] [-ca cert] [-ca-bundle bundle] \
                    [-ca-key key] [-int-bundle bundle] [-int-dir dir] [-port port] \
                    [-metadata file] [-remote remote_host] [-config config] \
                    [-responder cert] [-responder-key key] \
                    [-tls-cert cert] [-tls-key key] [-mutual-tls-ca ca] [-mutual-tls-cn regex] \
                    [-tls-remote-ca ca] [-mutual-tls-client-cert cert] [-mutual-tls-client-key key] \
                    [-db-config db-config] [-disable endpoint[,endpoint]]

Flags:
`

var serverFlags = []string{"address", "port", "min-tls-version", "ca", "ca-key", "ca-bundle", "int-bundle", "int-dir",
	"metadata", "remote", "config", "responder", "responder-key", "tls-key", "tls-cert", "mutual-tls-ca",
	"mutual-tls-cn", "tls-remote-ca", "mutual-tls-client-cert", "mutual-tls-client-key", "db-config", "disable"}

var (
	conf       cli.Config
	s          signer.Signer
	ocspSigner ocsp.Signer
	db         *sqlx.DB
)

var V1APIPrefix = "/api/v1/cfssl/"

func v1APIPath(path string) string { _ = "STUB: not implemented"; return "" }

//go:embed static
var staticContent embed.FS

var staticRedirections = map[string]string{
	"bundle":   "index.html",
	"scan":     "index.html",
	"packages": "index.html",
}

type staticFS struct {
	fs           fs.FS
	redirections map[string]string
}

func (s *staticFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

var errBadSigner = errors.New("signer not initialized")
var errNoCertDBConfigured = errors.New("cert db not configured (missing -db-config)")

var endpoints = map[string]func() (http.Handler, error){
	"sign": func() (http.Handler, error) {
		if s == nil {
			return nil, errBadSigner
		}

		h, err := signhandler.NewHandlerFromSigner(s)
		if err != nil {
			return nil, err
		}

		if conf.CABundleFile != "" && conf.IntBundleFile != "" {
			sh := h.Handler.(*signhandler.Handler)
			if err := sh.SetBundler(conf.CABundleFile, conf.IntBundleFile); err != nil {
				return nil, err
			}
		}

		return h, nil
	},

	"authsign": func() (http.Handler, error) {
		if s == nil {
			return nil, errBadSigner
		}

		h, err := signhandler.NewAuthHandlerFromSigner(s)
		if err != nil {
			return nil, err
		}

		if conf.CABundleFile != "" && conf.IntBundleFile != "" {
			sh := h.(*api.HTTPHandler).Handler.(*signhandler.AuthHandler)
			if err := sh.SetBundler(conf.CABundleFile, conf.IntBundleFile); err != nil {
				return nil, err
			}
		}

		return h, nil
	},

	"info": func() (http.Handler, error) {
		if s == nil {
			return nil, errBadSigner
		}
		return info.NewHandler(s)
	},

	"crl": func() (http.Handler, error) {
		if s == nil {
			return nil, errBadSigner
		}

		if db == nil {
			return nil, errNoCertDBConfigured
		}

		return crl.NewHandler(certsql.NewAccessor(db), conf.CAFile, conf.CAKeyFile)
	},

	"gencrl": func() (http.Handler, error) {
		if s == nil {
			return nil, errBadSigner
		}
		return gencrl.NewHandler(), nil
	},

	"newcert": func() (http.Handler, error) {
		if s == nil {
			return nil, errBadSigner
		}
		h := generator.NewCertGeneratorHandlerFromSigner(generator.CSRValidate, s)
		if conf.CABundleFile != "" && conf.IntBundleFile != "" {
			cg := h.(api.HTTPHandler).Handler.(*generator.CertGeneratorHandler)
			if err := cg.SetBundler(conf.CABundleFile, conf.IntBundleFile); err != nil {
				return nil, err
			}
		}
		return h, nil
	},

	"bundle": func() (http.Handler, error) {
		return bundle.NewHandler(conf.CABundleFile, conf.IntBundleFile)
	},

	"newkey": func() (http.Handler, error) {
		return generator.NewHandler(generator.CSRValidate)
	},

	"init_ca": func() (http.Handler, error) {
		return initca.NewHandler(), nil
	},

	"scan": func() (http.Handler, error) {
		return scan.NewHandler(conf.CABundleFile)
	},

	"scaninfo": func() (http.Handler, error) {
		return scan.NewInfoHandler(), nil
	},

	"certinfo": func() (http.Handler, error) {
		if db != nil {
			return certinfo.NewAccessorHandler(certsql.NewAccessor(db)), nil
		}

		return certinfo.NewHandler(), nil
	},

	"ocspsign": func() (http.Handler, error) {
		if ocspSigner == nil {
			return nil, errBadSigner
		}
		return apiocsp.NewHandler(ocspSigner), nil
	},

	"revoke": func() (http.Handler, error) {
		if db == nil {
			return nil, errNoCertDBConfigured
		}
		return revoke.NewHandler(certsql.NewAccessor(db)), nil
	},

	"/": func() (http.Handler, error) {
		subFS, _ := fs.Sub(staticContent, "static")
		return http.FileServer(http.FS(&staticFS{fs: subFS, redirections: staticRedirections})), nil
	},

	"health": func() (http.Handler, error) {
		return health.NewHealthCheck(), nil
	},

	"certadd": func() (http.Handler, error) {
		return certadd.NewHandler(certsql.NewAccessor(db), nil), nil
	},
}

func registerHandlers() { _ = "STUB: not implemented"; return }

func serverMain(args []string, c cli.Config) error { _ = "STUB: not implemented"; return nil }

var Command = &cli.Command{UsageText: serverUsageText, Flags: serverFlags, Main: serverMain}

var wrapHandler = defaultWrapHandler

func defaultWrapHandler(path string, handler http.Handler, err error) (string, http.Handler, error) {
	_ = "STUB: not implemented"
	return "", *new(http.Handler), nil
}

func SetWrapHandler(wh func(path string, handler http.Handler, err error) (string, http.Handler, error)) {
	_ = "STUB: not implemented"
	return
}

func SetEndpoint(path string, getHandler func() (http.Handler, error)) {
	_ = "STUB: not implemented"
	return
}

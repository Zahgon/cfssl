package exlib

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

var progname = filepath.Base(os.Args[0])

var Before = 5 * time.Minute

func Err(exit int, err error, format string, a ...interface{}) { _ = "STUB: not implemented"; return }

func Errx(exit int, format string, a ...interface{}) { _ = "STUB: not implemented"; return }

func Warn(err error, format string, a ...interface{}) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func Unpack(r io.Reader) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

const messageMax = 1 << 16

func Pack(w io.Writer, buf []byte) error { _ = "STUB: not implemented"; return nil }

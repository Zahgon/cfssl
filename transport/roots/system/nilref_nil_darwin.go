//go:build cgo && !arm && !arm64 && !ios && !go1.10
// +build cgo,!arm,!arm64,!ios,!go1.10

package system

/*
#cgo LDFLAGS: -framework CoreFoundation -framework Security
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"

func setNilCFRef(v *C.CFDataRef) { _ = "STUB: not implemented"; return }

func isNilCFRef(v C.CFDataRef) bool { _ = "STUB: not implemented"; return false }

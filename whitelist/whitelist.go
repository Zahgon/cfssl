package whitelist

import (
	"net"
	"sync"
)

type ACL interface {
	Permitted(net.IP) bool
}

type HostACL interface {
	ACL

	Add(net.IP)

	Remove(net.IP)
}

func validIP(ip net.IP) bool { _ = "STUB: not implemented"; return false }

type Basic struct {
	lock      *sync.Mutex
	whitelist map[string]bool
}

func (wl *Basic) Permitted(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func (wl *Basic) Add(ip net.IP) { _ = "STUB: not implemented"; return }

func (wl *Basic) Remove(ip net.IP) { _ = "STUB: not implemented"; return }

func NewBasic() *Basic { _ = "STUB: not implemented"; return nil }

func (wl *Basic) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (wl *Basic) UnmarshalJSON(in []byte) error { _ = "STUB: not implemented"; return nil }

func DumpBasic(wl *Basic) []byte { _ = "STUB: not implemented"; return nil }

func LoadBasic(in []byte) (*Basic, error) { _ = "STUB: not implemented"; return nil, nil }

type HostStub struct{}

func (wl HostStub) Permitted(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func (wl HostStub) Add(ip net.IP) { _ = "STUB: not implemented"; return }

func (wl HostStub) Remove(ip net.IP) { _ = "STUB: not implemented"; return }

func NewHostStub() HostStub { _ = "STUB: not implemented"; return *new(HostStub) }

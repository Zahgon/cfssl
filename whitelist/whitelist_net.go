package whitelist

import (
	"net"
	"sync"
)

type NetACL interface {
	ACL

	Add(*net.IPNet)

	Remove(*net.IPNet)
}

type BasicNet struct {
	lock      *sync.Mutex
	whitelist []*net.IPNet
}

func (wl *BasicNet) Permitted(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func (wl *BasicNet) Add(n *net.IPNet) { _ = "STUB: not implemented"; return }

func (wl *BasicNet) Remove(n *net.IPNet) { _ = "STUB: not implemented"; return }

func NewBasicNet() *BasicNet { _ = "STUB: not implemented"; return nil }

func (wl *BasicNet) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (wl *BasicNet) UnmarshalJSON(in []byte) error { _ = "STUB: not implemented"; return nil }

type NetStub struct{}

func (wl NetStub) Permitted(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func (wl NetStub) Add(ip *net.IPNet) { _ = "STUB: not implemented"; return }

func (wl NetStub) Remove(ip *net.IPNet) { _ = "STUB: not implemented"; return }

func NewNetStub() NetStub { _ = "STUB: not implemented"; return *new(NetStub) }

package auth

type AuthenticatedRequest struct {
	Timestamp     int64  `json:"timestamp,omitempty"`
	RemoteAddress []byte `json:"remote_address,omitempty"`
	Token         []byte `json:"token"`
	Request       []byte `json:"request"`
}

type Provider interface {
	Token(req []byte) (token []byte, err error)
	Verify(aReq *AuthenticatedRequest) bool
}

type Standard struct {
	key []byte
	ad  []byte
}

func New(key string, ad []byte) (*Standard, error) { _ = "STUB: not implemented"; return nil, nil }

func (p Standard) Token(req []byte) (token []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p Standard) Verify(ad *AuthenticatedRequest) bool { _ = "STUB: not implemented"; return false }

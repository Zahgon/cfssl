package tls

type clientHelloMsg struct {
	raw                 []byte
	vers                uint16
	random              []byte
	sessionId           []byte
	cipherSuites        []uint16
	compressionMethods  []uint8
	nextProtoNeg        bool
	serverName          string
	ocspStapling        bool
	scts                bool
	supportedCurves     []CurveID
	supportedPoints     []uint8
	ticketSupported     bool
	sessionTicket       []uint8
	signatureAndHashes  []signatureAndHash
	secureRenegotiation bool
	alpnProtocols       []string
}

func (m *clientHelloMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *clientHelloMsg) marshal() []byte { _ = "STUB: not implemented"; return nil }

func (m *clientHelloMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

type serverHelloMsg struct {
	raw                 []byte
	vers                uint16
	random              []byte
	sessionId           []byte
	cipherSuite         uint16
	compressionMethod   uint8
	nextProtoNeg        bool
	nextProtos          []string
	ocspStapling        bool
	scts                [][]byte
	ticketSupported     bool
	secureRenegotiation bool
	alpnProtocol        string
}

func (m *serverHelloMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *serverHelloMsg) marshal() []byte { _ = "STUB: not implemented"; return nil }

func (m *serverHelloMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

type certificateMsg struct {
	raw          []byte
	certificates [][]byte
}

func (m *certificateMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *certificateMsg) marshal() (x []byte) { _ = "STUB: not implemented"; return nil }

func (m *certificateMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

type serverKeyExchangeMsg struct {
	raw []byte
	key []byte
}

func (m *serverKeyExchangeMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *serverKeyExchangeMsg) marshal() []byte { _ = "STUB: not implemented"; return nil }

func (m *serverKeyExchangeMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

type certificateStatusMsg struct {
	raw        []byte
	statusType uint8
	response   []byte
}

func (m *certificateStatusMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *certificateStatusMsg) marshal() []byte { _ = "STUB: not implemented"; return nil }

func (m *certificateStatusMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

type serverHelloDoneMsg struct{}

func (m *serverHelloDoneMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *serverHelloDoneMsg) marshal() []byte { _ = "STUB: not implemented"; return nil }

func (m *serverHelloDoneMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

type clientKeyExchangeMsg struct {
	raw        []byte
	ciphertext []byte
}

func (m *clientKeyExchangeMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *clientKeyExchangeMsg) marshal() []byte { _ = "STUB: not implemented"; return nil }

func (m *clientKeyExchangeMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

type finishedMsg struct {
	raw        []byte
	verifyData []byte
}

func (m *finishedMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *finishedMsg) marshal() (x []byte) { _ = "STUB: not implemented"; return nil }

func (m *finishedMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

type nextProtoMsg struct {
	raw   []byte
	proto string
}

func (m *nextProtoMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *nextProtoMsg) marshal() []byte { _ = "STUB: not implemented"; return nil }

func (m *nextProtoMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

type certificateRequestMsg struct {
	raw []byte

	hasSignatureAndHash bool

	certificateTypes       []byte
	signatureAndHashes     []signatureAndHash
	certificateAuthorities [][]byte
}

func (m *certificateRequestMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *certificateRequestMsg) marshal() (x []byte) { _ = "STUB: not implemented"; return nil }

func (m *certificateRequestMsg) unmarshal(data []byte) bool {
	_ = "STUB: not implemented"
	return false
}

type certificateVerifyMsg struct {
	raw                 []byte
	hasSignatureAndHash bool
	signatureAndHash    signatureAndHash
	signature           []byte
}

func (m *certificateVerifyMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *certificateVerifyMsg) marshal() (x []byte) { _ = "STUB: not implemented"; return nil }

func (m *certificateVerifyMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

type newSessionTicketMsg struct {
	raw    []byte
	ticket []byte
}

func (m *newSessionTicketMsg) equal(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (m *newSessionTicketMsg) marshal() (x []byte) { _ = "STUB: not implemented"; return nil }

func (m *newSessionTicketMsg) unmarshal(data []byte) bool { _ = "STUB: not implemented"; return false }

func eqUint16s(x, y []uint16) bool { _ = "STUB: not implemented"; return false }

func eqCurveIDs(x, y []CurveID) bool { _ = "STUB: not implemented"; return false }

func eqStrings(x, y []string) bool { _ = "STUB: not implemented"; return false }

func eqByteSlices(x, y [][]byte) bool { _ = "STUB: not implemented"; return false }

func eqSignatureAndHashes(x, y []signatureAndHash) bool { _ = "STUB: not implemented"; return false }

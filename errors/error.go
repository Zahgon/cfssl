package errors

type Error struct {
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
}

type Category int

type Reason int

const (
	Success Category = 1000 * iota

	CertificateError

	PrivateKeyError

	IntermediatesError

	RootError

	PolicyError

	DialError

	APIClientError

	OCSPError

	CSRError

	CTError

	CertStoreError
)

const (
	None Reason = iota
)

const (
	BundleExpiringBit int = 1 << iota
	BundleNotUbiquitousBit
)

const (
	Unknown Reason = iota
	ReadFailed
	DecodeFailed
	ParseFailed
)

const (
	SelfSigned Reason = 100 * (iota + 1)

	VerifyFailed

	BadRequest

	MissingSerial
)

const (
	certificateInvalid = 10 * (iota + 1)
	unknownAuthority
)

const (
	Encrypted Reason = 100 * (iota + 1)

	NotRSAOrECCOrEd25519

	KeyMismatch

	GenerationFailed

	Unavailable
)

const (
	NoKeyUsages Reason = 100 * (iota + 1)

	InvalidPolicy

	InvalidRequest

	UnknownProfile

	UnmatchedWhitelist
)

const (
	AuthenticationFailure Reason = 100 * (iota + 1)

	JSONError

	IOError

	ClientHTTPError

	ServerRequestFailed
)

const (
	IssuerMismatch Reason = 100 * (iota + 1)

	InvalidStatus
)

const (
	PrecertSubmissionFailed = 100 * (iota + 1)

	CTClientConstructionFailed

	PrecertMissingPoison

	PrecertInvalidPoison
)

const (
	InsertionFailed = 100 * (iota + 1)

	RecordNotFound
)

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func New(category Category, reason Reason) *Error { _ = "STUB: not implemented"; return nil }

func Wrap(category Category, reason Reason, err error) *Error {
	_ = "STUB: not implemented"
	return nil
}

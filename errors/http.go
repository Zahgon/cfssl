package errors

type HTTPError struct {
	StatusCode int
	error
}

func (e *HTTPError) Error() string { _ = "STUB: not implemented"; return "" }

func NewMethodNotAllowed(method string) *HTTPError { _ = "STUB: not implemented"; return nil }

func NewBadRequest(err error) *HTTPError { _ = "STUB: not implemented"; return nil }

func NewBadRequestString(s string) *HTTPError { _ = "STUB: not implemented"; return nil }

func NewBadRequestMissingParameter(s string) *HTTPError { _ = "STUB: not implemented"; return nil }

func NewBadRequestUnwantedParameter(s string) *HTTPError { _ = "STUB: not implemented"; return nil }

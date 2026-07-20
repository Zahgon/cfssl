package api

import (
	"net/http"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}

type HTTPHandler struct {
	Handler
	Methods []string
}

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func (f HandlerFunc) Handle(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func HandleError(w http.ResponseWriter, err error) (code int) { _ = "STUB: not implemented"; return 0 }

func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func readRequestBlob(r *http.Request) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ProcessRequestOneOf(r *http.Request, keywordSets [][]string) (map[string]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ProcessRequestFirstMatchOf(r *http.Request, keywordSets [][]string) (map[string]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func matchKeywords(blob map[string]string, keywords []string) bool {
	_ = "STUB: not implemented"
	return false
}

type ResponseMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Success  bool              `json:"success"`
	Result   interface{}       `json:"result"`
	Errors   []ResponseMessage `json:"errors"`
	Messages []ResponseMessage `json:"messages"`
}

func NewSuccessResponse(result interface{}) Response {
	_ = "STUB: not implemented"
	return *new(Response)
}

func NewSuccessResponseWithMessage(result interface{}, message string, code int) Response {
	_ = "STUB: not implemented"
	return *new(Response)
}

func NewErrorResponse(message string, code int) Response {
	_ = "STUB: not implemented"
	return *new(Response)
}

func SendResponse(w http.ResponseWriter, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func SendResponseWithMessage(w http.ResponseWriter, result interface{}, message string, code int) error {
	_ = "STUB: not implemented"
	return nil
}

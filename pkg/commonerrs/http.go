package commonerrs

import "net/http"

// HTTPError provides some functional behavior to http errors
type HTTPError interface {
	GetStatusCode() int
	GetMessage() string
}

// httpError is used within elitecore as type implementing HTTPError
type httpError struct {
	StatusCode int
	Message    string
}

func NewHttpError(statusCode int, message string) HTTPError {
	return httpError{statusCode, message}
}

// GetStatusCode returns http status code of http error
func (e httpError) GetStatusCode() int {
	return e.StatusCode
}

// GetMessage returns err msg for http error
func (e httpError) GetMessage() string {
	return e.Message
}

var (
	// InternalServer returns standard InternalServerError as HTTPError
	InternalServer = httpError{http.StatusInternalServerError, "Internal Server Error"}
	// NotFound returns standard StatusNotFound as HTTPError
	NotFound = httpError{http.StatusNotFound, "Not found"}
	// BadRequest returns standard StatusBadRequest as HTTPError
	BadRequest = httpError{http.StatusBadRequest, "Bad request"}
	// Unauthorized returns standard StatusUnauthorized as HTTPError
	Unauthorized = httpError{http.StatusUnauthorized, "Unauthorized"}
	// Forbidden returns standard StatusForbidden as HTTPError
	Forbidden = httpError{http.StatusForbidden, "Forbidden"}
	// MethodNotAllowed returns standard StatusMethodNotAllowed as HTTPError
	MethodNotAllowed = httpError{http.StatusMethodNotAllowed, "Method not found"}
)

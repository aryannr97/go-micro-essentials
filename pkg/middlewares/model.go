package middlewares

import "net/http"

// ResponseRecorder is used to record the Status of the API response.
type ResponseRecorder struct {
	http.ResponseWriter
	Status int
}

// WriteHeader is added to ResponseRecorder to implement the http.ResponseWriter interface
func (r *ResponseRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

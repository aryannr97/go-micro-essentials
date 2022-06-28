package utils

import (
	"encoding/json"
	"net/http"

	"github.com/aryannr97/go-micro-essentials/pkg/commonerrs"
)

// DispatchSuccessJSON is used to send successful http server json response
func DispatchSuccessJSON(w http.ResponseWriter, data interface{}) {
	DispatchJSONWithStatus(w, http.StatusOK, data)
}

// DispatchNoContent is used to notify successful http server operation without response body
func DispatchNoContent(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// DispatchCreationJSON is used to send successful resource creation on http server
func DispatchCreationJSON(w http.ResponseWriter, data interface{}, resourceLoc string) {
	w.Header().Add("Location", resourceLoc)
	DispatchJSONWithStatus(w, http.StatusCreated, data)
}

// DispatchError is used to send http error response
func DispatchError(w http.ResponseWriter, err commonerrs.HTTPError) {
	DispatchJSONWithStatus(w, err.GetStatusCode(), err)
}

// DispatchJSONWithStatus prepares http json response with given data and status codes
func DispatchJSONWithStatus(w http.ResponseWriter, statusCode int, data interface{}) (int, error) {
	responseBytes, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return w.Write(responseBytes)
}

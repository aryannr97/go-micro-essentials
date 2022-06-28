package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aryannr97/go-micro-essentials/pkg/commonerrs"
	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}

func TestDispatchers(t *testing.T) {
	tests := []struct {
		basePath       string
		reqUri         string
		handler        http.HandlerFunc
		expectedStatus int
	}{
		{
			basePath:       "/test",
			reqUri:         "/test?type=success",
			handler:        testHandler,
			expectedStatus: http.StatusOK,
		},
		{
			basePath:       "/test",
			reqUri:         "/test?type=created",
			handler:        testHandler,
			expectedStatus: http.StatusCreated,
		},
		{
			basePath:       "/test",
			reqUri:         "/test?type=nocontent",
			handler:        testHandler,
			expectedStatus: http.StatusNoContent,
		},
		{
			basePath:       "/test",
			reqUri:         "/test?type=error",
			handler:        testHandler,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			basePath:       "/test",
			reqUri:         "/test?type=successwitherr",
			handler:        testHandler,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		r, _ := http.NewRequest(http.MethodGet, tt.reqUri, nil)
		w := httptest.NewRecorder()

		router := getRouter()
		router.Path(tt.basePath).Methods(http.MethodGet).HandlerFunc(tt.handler)
		router.ServeHTTP(w, r)

		if w.Code != tt.expectedStatus {
			t.Errorf("TestDispatchers() status got = %v, want = %v", w.Code, tt.expectedStatus)
		}
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Query().Get("type") {
	case "success":
		DispatchSuccessJSON(w, map[string]string{"status": "ok"})
	case "successwitherr":
		DispatchSuccessJSON(w, make(chan struct{}))
	case "created":
		DispatchCreationJSON(w, map[string]string{"status": "created"}, "http://host_url/path/resourceid")
	case "nocontent":
		DispatchNoContent(w)
	case "error":
		DispatchError(w, commonerrs.InternalServer)
	}
}

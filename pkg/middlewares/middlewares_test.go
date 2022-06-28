package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aryannr97/go-micro-essentials/pkg/logger"
	"github.com/aryannr97/go-micro-essentials/pkg/utils"
	"github.com/gorilla/mux"
)

func init() {
	setUpStandardLogger()
}

func setUpStandardLogger() {
	settings := logger.Settings{
		LogLevel: logger.Debug,
	}
	logger.Setup(settings)
}

func getRouter(middlware mux.MiddlewareFunc) *mux.Router {
	router := mux.NewRouter()
	router.Use(middlware)
	return router
}

func TestMiddlewares(t *testing.T) {
	tests := []struct {
		name           string
		apiEndpoint    string
		handler        http.HandlerFunc
		middleware     mux.MiddlewareFunc
		expectedStatus int
	}{
		{
			"LogRequest",
			"/test",
			testHandler,
			LogRequest,
			http.StatusOK,
		},
		{
			"PanicRecovery",
			"/panic",
			panicHandler,
			PanicRecovery,
			http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		r, _ := http.NewRequest(http.MethodGet, tt.apiEndpoint, nil)
		w := httptest.NewRecorder()

		router := getRouter(tt.middleware)
		router.Path(tt.apiEndpoint).Methods(http.MethodGet).HandlerFunc(tt.handler)
		router.ServeHTTP(w, r)

		if w.Code != tt.expectedStatus {
			t.Errorf("TestMiddlewares_%v status got = %v, want = %v", tt.name, w.Code, tt.expectedStatus)
		}
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	utils.DispatchSuccessJSON(w, map[string]string{"status": "ok"})
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("handler panic")
}

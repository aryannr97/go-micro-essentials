package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aryannr97/go-micro-essentials/pkg/commonerrs"
	"github.com/aryannr97/go-micro-essentials/pkg/logger"
	"github.com/aryannr97/go-micro-essentials/pkg/utils"
)

// LogRequest is a middleware for logging each http request made to server
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		r = r.WithContext(logger.FetchorPrepareCtx(r))

		recorder := &ResponseRecorder{
			ResponseWriter: w,
			Status:         200,
		}
		next.ServeHTTP(recorder, r)
		logger.RecordRequestDetails(r, recorder.Status, time.Since(start))
	})
}

// PanicRecovery handles panic on each http request
// It sends generic internal server error response back to client
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			ctx := logger.FetchorPrepareCtx(r)
			r = r.WithContext(ctx)

			if err := recover(); err != nil {
				logger.RecordError(ctx, fmt.Errorf("%v", err))
				utils.DispatchError(w, commonerrs.InternalServer)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

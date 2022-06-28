package logger

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// RecordRequestDetails logs details of http request made to server
func RecordRequestDetails(r *http.Request, status int, apiDuration time.Duration) {
	ctx := r.Context()
	entry := getMetadata(ctx).entry
	entry.WithFields(logrus.Fields{
		"caller": getCallerDetails(),
		"path":   r.URL.Path,
		"params": logrus.Fields{
			"path":  mux.Vars(r),
			"query": r.URL.Query(),
		},
		"method":       r.Method,
		"responseCode": status,
		"durationUs":   apiDuration.Microseconds(),
	}).Info()
}

// RecordInfo logs informative details within server
func RecordInfo(ctx context.Context, info interface{}) {
	entry := getMetadata(ctx).entry

	entry.WithFields(logrus.Fields{
		"caller":  getCallerDetails(),
		"details": info,
	}).Infof("Internal information")
}

// RecordError logs internal error occured within server
func RecordError(ctx context.Context, err error) {
	entry := getMetadata(ctx).entry

	entry.WithFields(logrus.Fields{
		"caller": getCallerDetails(),
	}).Errorf("%v", err)
}

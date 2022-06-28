package logger

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/aryannr97/go-micro-essentials/pkg/commonerrs"
	"github.com/google/uuid"
)

func setUpStandardLogger() {
	settings := Settings{
		LogLevel: Debug,
	}
	Setup(settings)
}

func setUpLoggerWithFile() *os.File {
	settings := Settings{
		LogLevel: Debug,
		LogFile:  "testlogs.txt",
	}
	return Setup(settings)
}

func getReqWitheHeader(r *http.Request, id string) *http.Request {
	r.Header.Set(RequestIDHeader, id)
	ctx := FetchorPrepareCtx(r)
	return r.WithContext(ctx)
}

func TestRequestDetails(t *testing.T) {
	setUpStandardLogger()
	request, _ := http.NewRequest(http.MethodGet, "/test", nil)
	type args struct {
		r           *http.Request
		status      int
		apiDuration time.Duration
	}
	tests := []struct {
		name      string
		requestId string
		args      args
	}{
		{
			"Log request details with passed request id",
			uuid.NewString(),
			args{
				r:           request,
				status:      http.StatusOK,
				apiDuration: 20 * time.Microsecond,
			},
		},
		{
			"Log request details with request id generated at server",
			"",
			args{
				r:           request,
				status:      http.StatusOK,
				apiDuration: 20 * time.Microsecond,
			},
		},
	}
	for _, tt := range tests {
		tt.args.r = getReqWitheHeader(tt.args.r, tt.requestId)

		t.Run(tt.name, func(t *testing.T) {
			RecordRequestDetails(tt.args.r, tt.args.status, tt.args.apiDuration)
		})
	}
}

func TestRecordInfo(t *testing.T) {
	setUpLoggerWithFile()
	ctx := context.WithValue(context.Background(), RequestID, uuid.New())
	type args struct {
		ctx  context.Context
		info interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Log internal info",
			args{
				ctx,
				map[string]interface{}{
					"customeridId": "cust123",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RecordInfo(tt.args.ctx, tt.args.info)
		})
	}

	os.Remove("testlogs.txt")
}

func TestRecordError(t *testing.T) {
	setUpStandardLogger()
	ctx := context.WithValue(context.Background(), RequestID, uuid.New())
	type args struct {
		ctx context.Context
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Log internal info",
			args{
				ctx,
				commonerrs.ErrDataMissing,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RecordError(tt.args.ctx, tt.args.err)
		})
	}
}

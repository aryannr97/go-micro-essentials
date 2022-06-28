package logger

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// metadata is used to persist essential data per context
type metadata struct {
	entry   *logrus.Entry
	request *http.Request
}

// callerDetails stores source file and func name of caller
type callerDetails struct {
	Name string `json:"name"`
	File string `json:"file"`
}

// Setup configures global settings for logging
func Setup(settings Settings) *os.File {
	logrus.SetLevel(logrus.Level(settings.LogLevel))
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if len(settings.LogFile) > 0 {
		fileHandler, err := os.OpenFile(settings.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Fatalf("Failed to setup log file %v: %v\n", settings.LogFile, err)
		}
		logrus.SetOutput(fileHandler)
		return fileHandler
	} else {
		logrus.SetOutput(os.Stdout)
	}

	return nil
}

// FetchorPrepareCtx either setup up request contex with additional attributes or return already set context
func FetchorPrepareCtx(r *http.Request) context.Context {
	ctx := r.Context()
	// Read request id from header if present else generate new one
	if id := r.Header.Get(RequestIDHeader); len(id) > 0 {
		ctx = context.WithValue(ctx, RequestID, requestKey(id))
	} else {
		ctx = context.WithValue(ctx, RequestID, uuid.New())
	}

	// Modify context with metadata to reuse as and when required
	if value := ctx.Value(ctxEntryKey); value == nil {
		entry := logrus.WithFields(logrus.Fields{"request_id": ctx.Value(RequestID)})
		m := &metadata{entry, r}
		ctx = context.WithValue(ctx, ctxEntryKey, m)
	}

	return ctx
}

// getMetadata returns new metadata or persisted within context
func getMetadata(ctx context.Context) *metadata {
	value := ctx.Value(ctxEntryKey)
	if value == nil {
		return &metadata{entry: logrus.NewEntry(logrus.StandardLogger())}
	}

	if md, ok := value.(*metadata); ok {
		return md
	}

	return &metadata{entry: logrus.NewEntry(logrus.StandardLogger())}
}

// getCallerDetails returns information about caller triggering logger methods
func getCallerDetails() *callerDetails {
	caller := &callerDetails{}
	if pc, file, line, ok := runtime.Caller(2); ok {
		funcName := runtime.FuncForPC(pc).Name()
		fn := strings.Split(funcName, "/")
		caller.Name = fn[len(fn)-1]

		fl := strings.Split(file, "/")
		fileName := fl[len(fl)-1]
		caller.File = fmt.Sprintf("%v:%v", fileName, line)
	}
	return caller
}

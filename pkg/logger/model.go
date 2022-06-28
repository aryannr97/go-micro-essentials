package logger

import "github.com/sirupsen/logrus"

// LogLevel wrapper type for logrus level
type LogLevel logrus.Level

// entry key is cutom type for setting in context
type entryKey int

// requestKey is custom type for request id
type requestKey string

const (
	Fatal           LogLevel   = LogLevel(logrus.FatalLevel)
	Warn            LogLevel   = LogLevel(logrus.WarnLevel)
	Debug           LogLevel   = LogLevel(logrus.DebugLevel)
	Info            LogLevel   = LogLevel(logrus.InfoLevel)
	Trace           LogLevel   = LogLevel(logrus.TraceLevel)
	RequestID       requestKey = "request_id"
	ctxEntryKey     entryKey   = 1
	RequestIDHeader string     = "X-Request-ID"
)

// Settings encapsulates all attributes required for configuring logger
type Settings struct {
	LogLevel LogLevel
	LogFile  string
}

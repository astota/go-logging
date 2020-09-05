package logging

import (
	"context"
	"sync"
)

var rMux sync.Mutex

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
	LevelFatal
)

// NewLogger return logger instance
func NewLogger() Logger {
	return newLogger()
}

// NewLoggerFunc type function. Each logger implementation must
// implement this and also this function need to create correctly
// initialized logger type (level and other configuration)
type NewLoggerFunc func() Logger

// newLogger contains function to create new logger instance
var newLogger NewLoggerFunc = newNilLogger

// loggers contains list of available logging methods
var loggers *map[string]NewLoggerFunc

// Logger is interface which logger must implement.
type Logger interface {
	// Debug level logging.
	Debug(string)
	// Debug level logging with formatting
	// Format string and its parameters
	Debugf(string, ...interface{})
	// Info level logging
	Info(string)
	// Info level logging with formatting
	// Format string and its parameters
	Infof(string, ...interface{})
	// Error level logging
	Error(string)
	// Error level logging with formatting
	// Format string and its parameters
	Errorf(string, ...interface{})
	// Fatal level logging
	Fatal(string)
	// Fatal level logging with formatting.
	// Format string and its parameters
	Fatalf(string, ...interface{})
	// Add field to logger. If logger does not suppor fields
	// this should be no-op call. All structured loggers
	// however should support this.
	// Fields are key value pairs which will be added to
	// structured log.
	AddFields(Fields) Logger
	// AddFieldToCurrent adds fields to current logger and return
	// also instance of itself. This is intended to use middleware
	// and similar places, where it is intended to add fields to current
	// logger.
	//
	// If logger does not suppor fields
	// this should be no-op call. All structured loggers
	// however should support this.
	// Fields are key value pairs which will be added to
	// structured log.
	AddFieldsToCurrent(Fields) Logger
	// Adds error content to logger
	WithError(error) Logger
	// SetLevel set logging level: debug,info,error,fatal
	SetLevel(Level)
}

// Fields presents fields which are added to logger
type Fields map[string]interface{}

// AddLogger adds new logger instance to logger pool
func Register(n string, l NewLoggerFunc) {
	rMux.Lock()
	if loggers == nil {
		l := make(map[string]NewLoggerFunc)
		loggers = &l
	}
	if _, ok := (*loggers)[n]; !ok {
		(*loggers)[n] = l
	}
	rMux.Unlock()
}

// UseLogger will change used backend
func UseLogger(n string) {
	rMux.Lock()
	if l, ok := (*loggers)[n]; ok {
		newLogger = l
	} else {
		newLogger = (*loggers)["nil"]
	}
	rMux.Unlock()
}

// init initializes package
func init() {
	Register("nil", newNilLogger)
}

// GetLogger will get logger from Context of create new
// instance if it does not exists
func GetLogger(ctx context.Context) Logger {
	var logger Logger
	loggerVal := ctx.Value("logger")
	if loggerVal == nil {
		logger = NewLogger()
	} else {
		var ok bool
		if logger, ok = loggerVal.(Logger); !ok {
			logger = NewLogger()
		}
	}
	return logger
}

// SetLogger sets given logger to context
func SetLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, "logger", l)
}

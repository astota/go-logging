package logging

import (
	"fmt"
	"reflect"
)

// TestlLogger is intended to use with tests as it can capture output
type TestLogger struct {
	TestOutput string
	InfoCount  int
	DebugCount int
	ErrorCount int
	FatalCount int
	Fields     map[string]interface{}
}

// ResetTestLogger resets global logger instance.
func ResetTestLogger() {
	testLogger = nil
}

func (l *TestLogger) Info(s string) {
	testLogger.TestOutput = fmt.Sprintf("%s%s", testLogger.TestOutput, s)
	testLogger.InfoCount++
}

func (l *TestLogger) Infof(f string, i ...interface{}) {
	testLogger.TestOutput = fmt.Sprintf("%s%s", testLogger.TestOutput, fmt.Sprintf(f, i...))
	testLogger.InfoCount++
}

func (l *TestLogger) Debug(s string) {
	testLogger.TestOutput = fmt.Sprintf("%s%s", testLogger.TestOutput, s)
	testLogger.DebugCount++
}

func (l *TestLogger) Debugf(f string, i ...interface{}) {
	testLogger.TestOutput = fmt.Sprintf("%s%s", testLogger.TestOutput, fmt.Sprintf(f, i...))
	testLogger.DebugCount++
}

func (l *TestLogger) Error(s string) {
	testLogger.TestOutput = fmt.Sprintf("%s%s", testLogger.TestOutput, s)
	testLogger.ErrorCount++
}

func (l *TestLogger) Errorf(f string, i ...interface{}) {
	testLogger.TestOutput = fmt.Sprintf("%s%s", testLogger.TestOutput, fmt.Sprintf(f, i...))
	testLogger.ErrorCount++
}

func (l *TestLogger) Fatal(s string) {
	testLogger.TestOutput = fmt.Sprintf("%s%s", testLogger.TestOutput, s)
	testLogger.FatalCount++
}

func (l *TestLogger) Fatalf(f string, i ...interface{}) {
	testLogger.TestOutput = fmt.Sprintf("%s%s", testLogger.TestOutput, fmt.Sprintf(f, i...))
	testLogger.FatalCount++
}

func (l *TestLogger) AddFields(f Fields) Logger {
	for k, v := range f {
		testLogger.Fields[k] = v
	}
	return testLogger
}

func (l *TestLogger) AddFieldsToCurrent(f Fields) Logger {
	for k, v := range f {
		testLogger.Fields[k] = v
	}
	return testLogger
}

func (l *TestLogger) WithError(err error) Logger {
	if err != nil {
		testLogger.Fields["error.message"] = err.Error()

		if t := reflect.TypeOf(err); t.Kind() == reflect.Ptr {
			testLogger.Fields["error.kind"] = t.Elem().Name()
		} else {
			testLogger.Fields["error.kind"] = t.Name()
		}
	}

	return testLogger
}

func (l *TestLogger) SetLevel(lvl Level) {
}

var testLogger *TestLogger

// newTestLogger creates TestLogger instance
func newTestLogger() Logger {
	if testLogger == nil {
		testLogger = &TestLogger{Fields: map[string]interface{}{}}
	}
	return testLogger
}

// RegisterTestLogger will register this test logger as "test-logger" to
// logger registry.
func RegisterTestLogger() {
	Register("test-logger", newTestLogger)
}

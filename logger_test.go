package logging

import (
	"context"
	"reflect"
	"testing"
)

func TestTryToGetInvalidLogger(t *testing.T) {
	UseLogger("1234567890")
	logger := NewLogger()
	if _, ok := logger.(*nilLogger); !ok {
		t.Errorf("NewLogger does not fallback to nil logger if invalid logger is used")
	}
}

func TestAddNewLogger(t *testing.T) {
	l := len(*loggers)
	Register("test1", newNilLogger)
	if l+1 != len(*loggers) {
		t.Errorf("Cannot add new logger")
	}
}

func TestDebugSimple(t *testing.T) {
	UseLogger("test-logger")
	logger := NewLogger().(*TestLogger)
	logger.Debugf("debug test  %s", "111")
	logger.Debug("222")
	if logger.TestOutput != "debug test  111222" {
		t.Errorf("Debug output is not correct '%s'", logger.TestOutput)
	}
}

func TestDebugGetLogger(t *testing.T) {
	UseLogger("nil")
	tests := []struct {
		name     string
		ctx      context.Context
		expected string
	}{
		{"No logger defined", context.Background(), "*logging.nilLogger"},
		{"With logger", context.WithValue(context.Background(), "logger", &TestLogger{}), "*logging.TestLogger"},
		{"With invalid logger type", context.WithValue(context.Background(), "logger", 1), "*logging.nilLogger"},
	}

	for _, tst := range tests {
		log := GetLogger(tst.ctx)
		typ := reflect.ValueOf(log).Type().String()
		if typ != tst.expected {
			t.Errorf("Invalid logger type, expected %s got %s",
				tst.expected, typ)
		}
	}
}

// Set "default-logger" as logger that is used
func ExampleUseLogger() {
	UseLogger("default-logger")
}

// Get logger context aware logger from Context
func ExampleGetLogger() {
	// ctx should come from context of request/job
	ctx := context.Background()

	logger := GetLogger(ctx)

	logger.Info("working logger")
}

// Get new logger without prefilled fields
func ExampleNewLogger() {
	logger := NewLogger()

	logger.Info("working logger")
}

// Formatting logged message
func ExampleLogger_formatting() {
	logger := NewLogger()
	value := 1

	logger.Infof("Value %d", value)
}

// Adding fields to structured output
func ExampleLogger_fields() {
	logger := NewLogger()

	logger = logger.AddFields(
		Fields{
			"value": 10,
		},
	)

	// This will include "value":10 key value pair in output
	logger.Info("value logged")
}

// Chaining logger functions
func ExampleLogger_chaining() {
	NewLogger().AddFields(
		Fields{
			"value": 10,
		},
	).Info("value logged")
}

func init() {
	RegisterTestLogger()
}

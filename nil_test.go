package logging

import (
	"testing"
)

func TestNilLogger(t *testing.T) {
	UseLogger("nil")
	logger := NewLogger()

	// Call methods, nothing should happen
	logger.Debug("")
	logger.Debugf("")
	logger.Info("")
	logger.Infof("")
	logger.Error("")
	logger.Errorf("")
	logger.Fatal("")
	logger.Fatalf("")
	logger.SetLevel(LevelDebug)
	logger.AddFields(Fields{"test": "test"})
	logger.AddFieldsToCurrent(Fields{"test": "test"})
}

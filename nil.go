package logging

// NilLogger is default logger, which does not output anything
type nilLogger struct{}

// Debug no-op implemenation
func (l *nilLogger) Debug(s string) {}

// Debugf no-op implemenation
func (l *nilLogger) Debugf(f string, i ...interface{}) {}

// Info no-op implemenation
func (l *nilLogger) Info(s string) {}

// Infof no-op implemenation
func (l *nilLogger) Infof(f string, i ...interface{}) {}

// Error no-op implemenation
func (l *nilLogger) Error(s string) {}

// Errorf no-op implemenation
func (l *nilLogger) Errorf(f string, i ...interface{}) {}

// Fatal no-op implemenation
func (l *nilLogger) Fatal(s string) {}

// Fatalf no-op implemenation
func (l *nilLogger) Fatalf(f string, i ...interface{}) {}

// AddFields no-op implemenation
func (l *nilLogger) AddFields(f Fields) Logger {
	return l
}

// AddFieldsToCurrent no-op implementation
func (l *nilLogger) AddFieldsToCurrent(Fields) Logger {
	return l
}

// WithError no-op implementation
func (l *nilLogger) WithError(err error) Logger {
	return l
}

// SetLevel no-op implementation
func (l *nilLogger) SetLevel(lvl Level) {}

// newNilLogger creates NilLogger instance
func newNilLogger() Logger {
	return &nilLogger{}
}

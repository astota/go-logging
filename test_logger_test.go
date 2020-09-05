package logging

import (
	"testing"
)

type fi struct {
	f    string
	vars []interface{}
}

func makeParams(args ...interface{}) []interface{} {
	list := make([]interface{}, len(args))
	for i, e := range args {
		list[i] = e
	}
	return list
}

func TestInfo(t *testing.T) {
	tests := []struct {
		name   string
		inputs []string
		output string
	}{
		{"one line", []string{"test line"}, "test line"},
		{"multiple lines", []string{"test line", "line two", "line three"}, "test lineline twoline three"},
	}

	for _, tst := range tests {
		UseLogger("test-logger")
		ResetTestLogger()
		logger := NewLogger()
		for _, line := range tst.inputs {
			logger.Info(line)
		}

		tl, _ := logger.(*TestLogger)
		if tst.output != tl.TestOutput {
			t.Errorf("%s: incorrect output, expected: '%s', got: '%s'", tst.name, tst.output, tl.TestOutput)
		}

		if len(tst.inputs) != tl.InfoCount {
			t.Errorf("%s: incorrect call count, expected: %d, got %d", tst.name, len(tst.inputs), tl.InfoCount)
		}
	}
}

func TestInfof(t *testing.T) {
	tests := []struct {
		name   string
		inputs []fi
		output string
	}{
		{"one line", []fi{fi{"formatted %s", makeParams("test line")}}, "formatted test line"},
		{"multiple lines", []fi{fi{"formatted %s", makeParams("test line")}, fi{"formatted1 %s", makeParams("test line1")}}, "formatted test lineformatted1 test line1"},
		{"integer", []fi{fi{"formatted %10.2d", makeParams(1)}}, "formatted         01"},
		{"float", []fi{fi{"formatted %10.2f", makeParams(123.123)}}, "formatted     123.12"},
	}

	for _, tst := range tests {
		UseLogger("test-logger")
		ResetTestLogger()
		logger := NewLogger()
		for _, line := range tst.inputs {
			logger.Infof(line.f, line.vars...)
		}

		tl, _ := logger.(*TestLogger)
		if tst.output != tl.TestOutput {
			t.Errorf("%s: incorrect output, expected: '%s', got: '%s'", tst.name, tst.output, tl.TestOutput)
		}

		if len(tst.inputs) != tl.InfoCount {
			t.Errorf("%s: incorrect call count, expected: %d, got %d", tst.name, len(tst.inputs), tl.InfoCount)
		}
	}
}

func TestDebug(t *testing.T) {
	tests := []struct {
		name   string
		inputs []string
		output string
	}{
		{"one line", []string{"test line"}, "test line"},
		{"multiple lines", []string{"test line", "line two", "line three"}, "test lineline twoline three"},
	}

	for _, tst := range tests {
		UseLogger("test-logger")
		ResetTestLogger()
		logger := NewLogger()
		for _, line := range tst.inputs {
			logger.Debug(line)
		}

		tl, _ := logger.(*TestLogger)
		if tst.output != tl.TestOutput {
			t.Errorf("%s: incorrect output, expected: '%s', got: '%s'", tst.name, tst.output, tl.TestOutput)
		}

		if len(tst.inputs) != tl.DebugCount {
			t.Errorf("%s: incorrect call count, expected: %d, got %d", tst.name, len(tst.inputs), tl.DebugCount)
		}
	}
}

func TestDebugf(t *testing.T) {
	tests := []struct {
		name   string
		inputs []fi
		output string
	}{
		{"one line", []fi{fi{"formatted %s", makeParams("test line")}}, "formatted test line"},
		{"multiple lines", []fi{fi{"formatted %s", makeParams("test line")}, fi{"formatted1 %s", makeParams("test line1")}}, "formatted test lineformatted1 test line1"},
		{"integer", []fi{fi{"formatted %10.2d", makeParams(1)}}, "formatted         01"},
		{"float", []fi{fi{"formatted %10.2f", makeParams(123.123)}}, "formatted     123.12"},
	}

	for _, tst := range tests {
		UseLogger("test-logger")
		ResetTestLogger()
		logger := NewLogger()
		for _, line := range tst.inputs {
			logger.Debugf(line.f, line.vars...)
		}

		tl, _ := logger.(*TestLogger)
		if tst.output != tl.TestOutput {
			t.Errorf("%s: incorrect output, expected: '%s', got: '%s'", tst.name, tst.output, tl.TestOutput)
		}

		if len(tst.inputs) != tl.DebugCount {
			t.Errorf("%s: incorrect call count, expected: %d, got %d", tst.name, len(tst.inputs), tl.DebugCount)
		}
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name   string
		inputs []string
		output string
	}{
		{"one line", []string{"test line"}, "test line"},
		{"multiple lines", []string{"test line", "line two", "line three"}, "test lineline twoline three"},
	}

	for _, tst := range tests {
		UseLogger("test-logger")
		ResetTestLogger()
		logger := NewLogger()
		for _, line := range tst.inputs {
			logger.Error(line)
		}

		tl, _ := logger.(*TestLogger)
		if tst.output != tl.TestOutput {
			t.Errorf("%s: incorrect output, expected: '%s', got: '%s'", tst.name, tst.output, tl.TestOutput)
		}

		if len(tst.inputs) != tl.ErrorCount {
			t.Errorf("%s: incorrect call count, expected: %d, got %d", tst.name, len(tst.inputs), tl.ErrorCount)
		}
	}
}

func TestErrorf(t *testing.T) {
	tests := []struct {
		name   string
		inputs []fi
		output string
	}{
		{"one line", []fi{fi{"formatted %s", makeParams("test line")}}, "formatted test line"},
		{"multiple lines", []fi{fi{"formatted %s", makeParams("test line")}, fi{"formatted1 %s", makeParams("test line1")}}, "formatted test lineformatted1 test line1"},
		{"integer", []fi{fi{"formatted %10.2d", makeParams(1)}}, "formatted         01"},
		{"float", []fi{fi{"formatted %10.2f", makeParams(123.123)}}, "formatted     123.12"},
	}

	for _, tst := range tests {
		UseLogger("test-logger")
		ResetTestLogger()
		logger := NewLogger()
		for _, line := range tst.inputs {
			logger.Errorf(line.f, line.vars...)
		}

		tl, _ := logger.(*TestLogger)
		if tst.output != tl.TestOutput {
			t.Errorf("%s: incorrect output, expected: '%s', got: '%s'", tst.name, tst.output, tl.TestOutput)
		}

		if len(tst.inputs) != tl.ErrorCount {
			t.Errorf("%s: incorrect call count, expected: %d, got %d", tst.name, len(tst.inputs), tl.ErrorCount)
		}
	}
}

func TestFatal(t *testing.T) {
	tests := []struct {
		name   string
		inputs []string
		output string
	}{
		{"one line", []string{"test line"}, "test line"},
		{"multiple lines", []string{"test line", "line two", "line three"}, "test lineline twoline three"},
	}

	for _, tst := range tests {
		UseLogger("test-logger")
		ResetTestLogger()
		logger := NewLogger()
		for _, line := range tst.inputs {
			logger.Fatal(line)
		}

		tl, _ := logger.(*TestLogger)
		if tst.output != tl.TestOutput {
			t.Errorf("%s: incorrect output, expected: '%s', got: '%s'", tst.name, tst.output, tl.TestOutput)
		}

		if len(tst.inputs) != tl.FatalCount {
			t.Errorf("%s: incorrect call count, expected: %d, got %d", tst.name, len(tst.inputs), tl.FatalCount)
		}
	}
}

func TestFatalf(t *testing.T) {
	tests := []struct {
		name   string
		inputs []fi
		output string
	}{
		{"one line", []fi{fi{"formatted %s", makeParams("test line")}}, "formatted test line"},
		{"multiple lines", []fi{fi{"formatted %s", makeParams("test line")}, fi{"formatted1 %s", makeParams("test line1")}}, "formatted test lineformatted1 test line1"},
		{"integer", []fi{fi{"formatted %10.2d", makeParams(1)}}, "formatted         01"},
		{"float", []fi{fi{"formatted %10.2f", makeParams(123.123)}}, "formatted     123.12"},
	}

	for _, tst := range tests {
		UseLogger("test-logger")
		ResetTestLogger()
		logger := NewLogger()
		for _, line := range tst.inputs {
			logger.Fatalf(line.f, line.vars...)
		}

		tl, _ := logger.(*TestLogger)
		if tst.output != tl.TestOutput {
			t.Errorf("%s: incorrect output, expected: '%s', got: '%s'", tst.name, tst.output, tl.TestOutput)
		}

		if len(tst.inputs) != tl.FatalCount {
			t.Errorf("%s: incorrect call count, expected: %d, got %d", tst.name, len(tst.inputs), tl.FatalCount)
		}
	}
}

func TestFields(t *testing.T) {
	tests := []struct {
		name   string
		fields map[string]interface{}
	}{
		{"no fields", map[string]interface{}{}},
		{"one field", map[string]interface{}{"one": 1}},
	}

	for _, tst := range tests {
		UseLogger("test-logger")
		ResetTestLogger()
		logger := NewLogger()
		logger = logger.AddFields(tst.fields)

		tl, _ := logger.(*TestLogger)

		if len(tl.Fields) != len(tst.fields) {
			t.Errorf("%s: incorrect number or fields: expected %d, got %d", tst.name, len(tst.fields), len(tl.Fields))
		}

		for k, _ := range tst.fields {
			if _, exists := tl.Fields[k]; !exists {
				t.Errorf("%s: Field '%s' missing", tst.name, k)
			}
		}

	}
}

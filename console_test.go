package console

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TryConsoleAt(level LEVEL) string {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	csl := NewConsole(INFO)
	csl.SetLevel(level)

	csl.Trace("trace-")
	csl.Debug("debug-")
	csl.Warn("warn-")
	csl.Info("info-")
	csl.Print("print")

	w.Close()
	os.Stdout = stdout

	out, _ := io.ReadAll(r)

	return string(out)
}

func TestConsole(t *testing.T) {
	result := TryConsoleAt(INFO)
	assert.Equal(t, "info-print", result)
	result = TryConsoleAt(WARN)
	assert.Equal(t, "warn-info-print", result)
	result = TryConsoleAt(DEBUG)
	assert.Equal(t, "debug-warn-info-print", result)
	result = TryConsoleAt(TRACE)
	assert.Equal(t, "trace-debug-warn-info-print", result)
}

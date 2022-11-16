package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithFields(t *testing.T) {
	fields := map[string]interface{}{
		"foo":    "bar",
		"number": 1,
	}

	opts := Options{}
	WithFields(fields)(&opts)

	assert.Equal(t, "bar", opts.Fields["foo"])
	assert.Equal(t, 1, opts.Fields["number"])

}

func TestWithConsole(t *testing.T) {
	opts := Options{}
	WithConsole(true)(&opts)

	assert.Equal(t, true, opts.ConsolePrint)
}

func TestWithLevel(t *testing.T) {
	opts := Options{}
	WithLevel(DebugLevel)(&opts)

	assert.Equal(t, DebugLevel, opts.Level)
}

func TestWithOutput(t *testing.T) {
	opts := Options{}
	WithOutput(os.Stdout)(&opts)

	assert.Equal(t, os.Stdout, opts.Out)
}

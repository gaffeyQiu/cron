package log

import "testing"

func TestLevelString(t *testing.T) {
	tests := []struct {
		input  Level
		expect string
	}{
		{DebugLevel, "debug"},
		{InfoLevel, "info"},
		{WarnLevel, "warn"},
		{ErrorLevel, "error"},
		{FatalLevel, "fatal"},
	}

	for _, tt := range tests {
		actual := (tt.input).String()
		if actual != tt.expect {
			t.Fatalf("input: %v, expect: %s, actual: %s", tt.input, tt.expect, actual)
		}
	}
}

func TestGetLevel(t *testing.T) {
	tests := []struct {
		input  string
		expect Level
	}{
		{"debug", DebugLevel},
		{"info", InfoLevel},
		{"warn", WarnLevel},
		{"error", ErrorLevel},
		{"fatal", FatalLevel},
	}

	for _, tt := range tests {
		actual, err := GetLevel(tt.input)
		if err != nil {
			t.Fail()
		}
		if actual != tt.expect {
			t.Fatalf("input: %v, expect: %s, actual: %s", tt.input, tt.expect, actual)
		}
	}
}

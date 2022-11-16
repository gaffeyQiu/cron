package helper

import (
	"fmt"
	"testing"
)

type testStringer struct {
	name string
}

func (str testStringer) String() string {
	return fmt.Sprintf("name: %s", str.name)
}

func TestMustStringVar(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect string
	}{
		{1, "1"}, // int -> string
		{-1, "-1"},
		{3.5, "3.5"}, // float -> string
		{-3.5, "-3.5"},
		{true, "true"}, // bool -> string
		{false, "false"},
		{"foo", "foo"},                           // string -> string
		{testStringer{name: "bar"}, "name: bar"}, // stringer -> string
	}

	// a := testStringer{name: "hello"}
	// fmt.Println(a.(fmt.Stringer))

	for _, tt := range tests {
		actual := MustStringVar(tt.input)
		if actual != tt.expect {
			t.Fatalf("input: %v, expect: %s, actual: %s", tt.input, tt.expect, actual)
		}
	}
}

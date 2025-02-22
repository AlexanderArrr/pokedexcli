package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Always on my mind ",
			expected: []string{"always", "on", "my", "mind"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("FAIL: Different length expected!\nExpected: %v\nLength of %d\nActual: %v\nLength of %d", c.expected, len(c.expected), actual, len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("FAIL: Strings are not matching!")
			}
		}
	}
}

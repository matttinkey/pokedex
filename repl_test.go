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
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "test",
			expected: []string{"test"},
		},
		{
			input:    "a FEW  more  ",
			expected: []string{"a", "few", "more"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("len_expected (%v) != len_actual (%v)", c.expected, actual)
			continue
		}

		for i := range c.expected {
			expectedWord := c.expected[i]
			if expectedWord == actual[i] {
				continue
			}
			t.Errorf("%v does not match %v", actual[i], expectedWord)
		}
	}
}

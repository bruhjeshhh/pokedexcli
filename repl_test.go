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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  you are already there  ",
			expected: []string{"you", "are", "already", "there"},
		},
		{
			input:    "  there's a lady who's sure  ",
			expected: []string{"there's", "a", "lady", "who's", "sure"},
		},
		{
			input:    "  all that glitters is gold  ",
			expected: []string{"all", "that", "glitters", "is", "gold"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Fail hogya")

		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("fail")

			}
		}
	}
}

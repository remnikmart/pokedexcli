package funcs

import (
	"fmt"
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
			input:    "one_word",
			expected: []string{"one_word"},
		},
		{
			input:    "woRDs IN diFFereNt Case",
			expected: []string{"words", "in", "different", "case"},
		},
		{
			input:    "check   for  emtpy     occurrences",
			expected: []string{"check", "for", "emtpy", "occurrences"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "  \n",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := CleanInput(c.input)
		for i := range c.expected {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				fmt.Printf("Input:    %s\n", c.input)
				fmt.Printf("Expected: %v\n", c.expected)
				fmt.Printf("Actual:   %v\n", actual)
				t.Errorf("%s is not equal %s\n", word, expectedWord)
			}
		}
	}
}

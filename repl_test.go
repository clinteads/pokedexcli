package main

import(
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
	}
	for _,c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths of actual %v and expected %v do not match",len(actual),len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("got %v, expected %v", word, expectedWord)
			}	
		}
	}
}
package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"Hello World", []string{"hello", "world"}},
		{"  Leading and trailing spaces  ", []string{"leading", "and", "trailing", "spaces"}},
		{"Mixed CASE Words", []string{"mixed", "case", "words"}},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Expected %d words, got %d for input '%s'", len(cs.expected), len(actual), cs.input)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("Expected word '%s', got '%s' for input '%s'", expectedWord, actualWord, cs.input)
			}
		}
	}
}

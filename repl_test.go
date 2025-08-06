package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	result := cleanInput("hello world")
	expected := []string{"hello", "world"}

	if len(result) != len(expected) {
		t.Errorf("lengths don't match. result=%v, expected=%v", len(result), len(expected))
	}

	for i, actualWord := range result {
		expectedWord := expected[i]
		if actualWord != expectedWord {
			t.Errorf("words don't match. result=%v, expected=%v", actualWord, expectedWord)
		}

	}
	// cases := []struct {
	// 	input    string
	// 	expected []string
	// }{
	// 	{
	// 		input:    " hello world ",
	// 		expected: []string{"hello", "world"},
	// 	},
	// }
	// for _, c := range cases {
	// 	actual := cleanInput(c.input)
	// 	if len(actual) != len(c.expected) {
	// 		t.Errorf("lengths don't match. expected=%v, actual=%v", len(c.expected), len(actual))
	// 		continue
	// 	}
	// 	for i := range actual {
	// 		word := actual[i]
	// 		expectedWord := c.expected[i]
	// 		if word != expectedWord {
	// 			t.Errorf("word don't match. expected=%v, actual=%v", expectedWord, word)
	// 		}
	// 	}
	// }
}

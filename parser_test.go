package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"", nil},
		{"foo", []string{"foo"}},
		{"foo bar", []string{"foo", "bar"}},
		{"'foo", []string{"'foo"}},
		{"foo'", []string{"foo'"}},
		{"foo'bar", []string{"foo'bar"}},
		{
			input:    "  That thereby beauty's rose might never die,",
			expected: []string{"That", "thereby", "beauty's", "rose", "might", "never", "die"},
		},
		{
			input:    "  Feed'st thy light's flame with self-substantial fuel,",
			expected: []string{"Feed'st", "thy", "light's", "flame", "with", "self-substantial", "fuel"},
		},
		{
			input:    "    Which used lives th' executor to be.",
			expected: []string{"Which", "used", "lives", "th'", "executor", "to", "be"},
		},
		{
			input:    "  But yield them up where I myself must render-",
			expected: []string{"But", "yield", "them", "up", "where", "I", "myself", "must", "render"},
		},
		{
			input:    "  '\"O then advance of yours that phraseless hand",
			expected: []string{"O", "then", "advance", "of", "yours", "that", "phraseless", "hand"},
		},
	}
	for i, testCase := range testCases {
		got := Parse(testCase.input)
		expected := testCase.expected
		gotStr := fmt.Sprintf("%v", got)
		expectedStr := fmt.Sprintf("%v", expected)
		if gotStr != expectedStr {
			t.Errorf("%d: got %q but expected %q", i, gotStr, expectedStr)
		}
	}
}

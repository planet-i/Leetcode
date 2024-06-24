package main

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"babad", "bab"}, // "aba" is also a valid answer
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"}, // "c" is also a valid answer
		{"racecar", "racecar"},
		{"anana", "anana"},
		{"", ""},
		{"abcd", "a"}, // Any single character is a valid answer
	}

	for _, tc := range testCases {
		result := longestPalindrome(tc.input)
		if result == tc.expected {
			fmt.Printf("PASS: Input: %s, Output: %s\n", tc.input, result)
		} else {
			t.Errorf("FAIL: Input: %s, Output: %s, Expected: %s\n", tc.input, result, tc.expected)
		}
	}
}

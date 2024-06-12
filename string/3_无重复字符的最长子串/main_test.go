package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},  // 因为无重复字符的最长子串是 "b"，所以其长度为 1。
		{"pwwkew", 3}, // 因为无重复字符的最长子串是 "wke"，所以其长度为 3。 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
		{"", 0},
		{"aab", 2},
		{"dvdf", 3},
	}

	for _, test := range tests {
		got := lengthOfLongestSubstring(test.input)
		if got != test.want {
			t.Errorf("lengthOfLongestSubstring(%s) = %d; want %d", test.input, got, test.want)
		}
	}
}

package main

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		expect string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"ab", "A", ""},
		{"a", "b", ""},
		{"aa", "aa", "aa"},
		{"aaflslflsldkalskaaa", "aaa", "aaa"},
	}

	for _, tt := range tests {
		result := minWindow(tt.s, tt.t)
		if result != tt.expect {
			t.Errorf("minWindow(%s, %s) = %s; expected %s", tt.s, tt.t, result, tt.expect)
		}
	}
}

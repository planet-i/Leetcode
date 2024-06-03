package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, test := range tests {
		result := twoSum1(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expect) {
			fmt.Printf("Test failed for input nums=%v, target=%d. Expected: %v, Got: %v\n", test.nums, test.target, test.expect, result)
		} else {
			fmt.Printf("Test passed for input nums=%v, target=%d\n", test.nums, test.target)
		}
	}
}

func TestTwoSum2(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, test := range tests {
		result := twoSum2(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expect) {
			fmt.Printf("Test failed for input nums=%v, target=%d. Expected: %v, Got: %v\n", test.nums, test.target, test.expect, result)
		} else {
			fmt.Printf("Test passed for input nums=%v, target=%d\n", test.nums, test.target)
		}
	}
}

func TestTwoSum3(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{2, 7}}, // 2 + 7 = 9
		{[]int{3, 2, 4}, 6, []int{2, 4}},      // 2 + 4 = 6
		{[]int{3, 3}, 6, []int{3, 3}},         // 3 + 3 = 6
	}

	for _, tt := range tests {
		got := twoSum3(tt.nums, tt.target)
		// 验证结果
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("twoSum3(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

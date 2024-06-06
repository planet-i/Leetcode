package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{[]int{1, 2, 3, 4, 5}, 20, [][]int{}}, // 无法通过数组中的任何四个数相加得到目标值
		{[]int{1, 2, 3}, 6, [][]int{}},        // 数组长度不足四个元素
		{[]int{}, 0, [][]int{}},               // 空数组
		{[]int{-1, 0, 1, 2, -1, -4}, -1, [][]int{{-4, 0, 1, 2}, {-1, -1, 0, 1}}},
	}

	for _, test := range tests {
		result := fourSum2(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expect) {
			fmt.Printf("Test failed for input nums=%v, target=%d. Expected: %v, Got: %v\n", test.nums, test.target, test.expect, result)
		} else {
			fmt.Printf("Test passed for input nums=%v, target=%d. Expected: %v, Got: %v\n", test.nums, test.target, test.expect, result)
		}
	}
}

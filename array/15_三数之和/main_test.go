package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum1(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect [][]int
	}{
		{
			name:   "Example case",
			nums:   []int{-1, 0, 1, 2, -1, -4},
			expect: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:   "No solution",
			nums:   []int{1, 2, 3, 4, 5},
			expect: [][]int{},
		},
		{
			name:   "Duplicates",
			nums:   []int{0, 0, 0, 0},
			expect: [][]int{{0, 0, 0}},
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			result := threeSum1(test.nums)
			// 注意：因为结果集是无序的，所以我们需要先对结果集中的每个子数组进行排序
			// 然后再比较是否相等
			for i := range result {
				sort.Ints(result[i])
			}
			if !reflect.DeepEqual(result, test.expect) {
				fmt.Printf("Test failed for input nums=%v. Expected: %v, Got: %v\n", test.nums, test.expect, result)
			} else {
				fmt.Printf("Test passed for input nums=%v\n", test.nums)
			}
		})
	}
}

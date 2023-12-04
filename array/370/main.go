// 370 差分数组
package main

import "fmt"

func main() {
	kk := [][]int{
		{1, 3, 2},
		{2, 4, 3},
		{0, 2, -2},
	}
	s := getModifiedArray(5, kk)
	fmt.Println(s)
}
func getModifiedArray(n int, update [][]int) []int {
	nums := make([]int, n)
	df := Constructor(nums)
	fmt.Println("---", df)
	for _, v := range update {
		i := v[0]
		j := v[1]
		inc := v[2]
		df = increment(i, j, inc, df)
		fmt.Println("-1-", df)
	}
	nums = Restore(df)
	return nums
}

func Constructor(nums []int) []int {
	n := len(nums)
	preDiff := make([]int, n)
	preDiff[0] = nums[0]
	for i := 1; i < n; i++ {
		preDiff[i] = nums[i] - nums[i-1]
	}
	return preDiff
}

func increment(i, j, val int, diff []int) []int {
	diff[i] += val
	if j+1 < len(diff) {
		diff[j+1] -= val
	}
	return diff
}

func Restore(diff []int) []int {
	n := len(diff)
	nums := make([]int, n)
	nums[0] = diff[0]
	for i := 1; i < n; i++ {
		nums[i] = nums[i-1] + diff[i]
	}
	return nums
}

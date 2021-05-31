package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	result := moveZeroes(nums)
	fmt.Println(result)
}

func moveZeroes(nums []int) []int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if k != i {
				nums[k], nums[i] = nums[i], nums[k]
				k++
			} else {
				k++
			}
		}
	}
	return nums
}

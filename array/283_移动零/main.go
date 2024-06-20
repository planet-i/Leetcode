/*
	283 移动零

题目：给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/
package main

func moveZeroes(nums []int) []int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if k != i {
				nums[k], nums[i] = nums[i], nums[k]
			}
			k++
		}
	}
	return nums
}

func moveZeroes2(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return
}

func moveZeroes3(nums []int) {
	// 去除 nums 中的所有 0
	// 返回去除 0 之后的数组长度
	p := removeElement(nums, 0)
	// 将 p 之后的所有元素赋值为 0
	for ; p < len(nums); p++ {
		nums[p] = 0
	}
}

// 双指针技巧，复用 [27. 移除元素] 的解法。
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

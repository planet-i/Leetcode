/*
27. 移除元素
题目：给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：
更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
返回 k。
*/
package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			// 先给 nums[slow] 赋值然后再 slow++，保证 nums[0..slow-1] 是不包含值为 val 的元素的
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	// slow索引前的位置都是不包含值为 val 的元素的
	return slow
}

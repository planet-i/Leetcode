/* 26. 删除有序数组中的重复项
题目：给你一个非严格递增排列的数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。元素的相对顺序应该保持 一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。返回 k 。
*/

package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0 // 定义快慢指针
	for fast < len(nums) {
		if nums[fast] != nums[slow] { // 快指针遇到了与慢指针所在位置不一样的数
			slow++                  // 将慢指针右移
			nums[slow] = nums[fast] // 慢指针位置的数被快指针位置的数替换掉
		}
		fast++ // 快指针一直加，直到遇到与慢指针所在位置不一样的数
	}
	return slow + 1
}

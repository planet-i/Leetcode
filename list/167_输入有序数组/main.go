/*
	167 两数之和2_输入有序数组

题目: 给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。
你可以假设每个输入 只对应唯一的答案 ，而且你不可以重复使用相同的元素。
你所设计的解决方案必须只使用常量级的额外空间。

注意：有序数组、返回的是索引、任意顺序、只返回一个序列
*/
package main

// 双指针法：要求返回索引序列，输入的数组有序，那就不担心排序后会影响索引。
// 时间复杂度：O(n)，其中 n 是数组的长度。两个指针移动的总次数最多为 n次。
// 空间复杂度：O(1)。
func twoSum1(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

// 二分查找法：固定第一个数，然后寻找第二个数
// 时间复杂度：O(N*logN)，其中 n是数组的长度。需要遍历数组一次确定第一个数，时间复杂度是 O(N)，寻找第二个数使用二分查找，时间复杂度是 O(logN)，因此总时间复杂度是  O(N*logN)。
// 空间复杂度：O(1)。
func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := (high-low)/2 + low
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return nil
}

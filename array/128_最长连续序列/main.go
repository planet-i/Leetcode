/*
	128 最长连续序列

题目：给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
注意：请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/
package main

// 哈希表：用哈希表来判断元素是否存在
// 时间复杂度：O(n)，其中 n 为数组的长度。
// 空间复杂度：O(n)。哈希表存储数组中所有的数需要 O(n) 的空间
func longestConsecutive(nums []int) int {
	// 转化成哈希集合，方便快速查找是否存在某个元素
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	// res 记录找到的连续序列的最大值
	res := 0
	for num := range set {
		if set[num-1] {
			// num 不是连续子序列的第一个，跳过
			continue
		}
		// num 是连续子序列的第一个，开始向上计算连续子序列的长度
		curNum := num
		curLen := 1

		for set[curNum+1] {
			curNum += 1
			curLen += 1
		}
		// 更新最长连续序列的长度
		res = max(res, curLen)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

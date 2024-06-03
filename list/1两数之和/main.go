/* 1 两数之和
题目: 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
     你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
     你可以按任意顺序返回答案。
注意：返回的是索引、任意顺序、只返回一个序列
*/

package main

import (
	"sort"
)

// twoSum1 暴力枚举法。枚举数组中的每一个数 x，寻找数组中是否存在 target - x。
// 时间复杂度：O(N^2) 其中 N 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
// 空间复杂度：O(1)。
func twoSum1(nums []int, target int) []int {
	// 边界条件判断
	n := len(nums)
	if n < 2 {
		return []int{}
	}
	// 遍历每个数
	for i := 0; i < n; i++ {
		// 将每个数与后面的数尝试相加一下，符合要求就立马返回
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				res := []int{i, j}
				return res
			}
		}
	}
	return nil
}

// twoSum2 哈希表辅助判断法。数组其实可以理解为一个「索引 -> 值」的哈希表映射，我们再建立一个「值 -> 索引」的映射，在找到值的时候获取索引。
// 时间复杂度：O(N)，其中 N 是数组中的元素数量。对于每一个元素 x，我们可以 O(1) 地寻找 target - x。
// 空间复杂度：O(N)，其中 N 是数组中的元素数量。主要为哈希表的开销。
func twoSum2(nums []int, target int) []int {
	valToIndex := make(map[int]int)
	for i, num := range nums {
		// 查表：对于一个元素 nums[i]，要找的数就是 target - nums[i]，判断哈希表中这个数是否存在
		need := target - num
		if index, ok := valToIndex[need]; ok {
			// 确定元素在map中存在，就取出元素索引
			return []int{index, i}
		}
		// 存入 val -> index 的映射
		valToIndex[num] = i
	}
	return nil
}

// twoSum3  返回元素，而不是索引的情况。数组排序后，双指针法，让两个指针从两端相向而行,指针相遇就结束。
func twoSum3(nums []int, target int) []int {
	// TODO 如果要求返回元素的索引，而排序会破坏元素的原始索引，需要先记录值和原始索引的映射。
	// 先对数组进行排序
	sort.Ints(nums)
	// 定义左右指针
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			// 让和大一点
			left++
		} else if sum > target {
			// 让和小一点
			right--
		} else if sum == target {
			// 找到两个数，返回
			return []int{nums[left], nums[right]}
		}
	}
	return nil
}

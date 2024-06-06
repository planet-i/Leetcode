/* 15 三数之和
题目: 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
	  同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

package main

import "sort"

// 双指针法：使用双指针来遍历数组，同时固定一个数作为基准，然后在剩余的数组中使用双指针找到两个数与基准数之和等于目标值的情况。在遍历过程中，需要注意去重。
// 算法的时间复杂度为 O(N^2)，其中 N 是数组的长度。
func threeSum1(nums []int) [][]int {
	var res [][]int
	// 先对数组进行排序
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		// 跳过重复的基准数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 定义这一轮的左右指针
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 跳过重复的左指针和右指针
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res
}

// 最坏的情况下，数组中的元素全部为 0
// 算法使用了两个指针 left 和 right 在排好序的数组中寻找满足条件的三元组。
// 外层循环遍历了数组，内层循环使用了双指针技术。 对于每个基准元素，内层循环的时间复杂度是 O(n)，
// 因为 left 和 right 指针最多会移动 n 次。由于外层循环遍历了数组中的每个元素，因此整个双指针遍历的时间复杂度为 O(n^2)。

// 暴力枚举法：这个方法像双指针写法的低配版
func threeSum2(nums []int) [][]int {
	n := len(nums)
	// 先排序、方便解决重复问题
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for a := 0; a < n; a++ {
		// a需要和上一次枚举的数不相同，重复的数需跳过
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		c := n - 1
		target := -1 * nums[a]
		// 枚举 b
		for b := a + 1; b < n; b++ {
			// b需要和上一次枚举的数不相同，重复的数需跳过
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for b < c && nums[b]+nums[c] > target {
				c--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if b == c {
				break
			}
			if nums[b]+nums[c] == target {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return ans
}

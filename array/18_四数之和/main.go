/*
	18 四数之和

题目: 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]]
注意：若两个四元组元素一一对应，则认为两个四元组重复。可以按 任意顺序 返回答案。a、b、c 和 d 互不相同。
*/
package main

import "sort"

// 双指针法：借鉴三数之和的思路，四数之和只是比它多嵌套一层循环
// 使用两重循环分别枚举前两个数，然后在两重循环枚举到的数之后使用双指针枚举剩下的两个数。
// 时间复杂度：O(N^3)
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left+1] == nums[left] {
						left++
					}
					for left < right && nums[right-1] == nums[right] {
						right--
					}
					left++
					right--
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return res
}

// 排序 + 双指针：四重循环枚举所有的四元组,使用哈希表进行去重操作 ---> 先排序，利用相邻判断是否相等去重。用双指针法去掉一重循环。添加一些剪枝操作。
// 时间复杂度：O(N^3)，其中 N 是数组的长度。排序的时间复杂度是 O(NlogN)，枚举四元组的时间复杂度是 O(N^3)，因此总时间复杂度为  O(N^3 + NlogN) = O(N^3)
func fourSum2(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	// 在确定第一个数之后，如果 nums[i]+nums[i+1]+nums[i+2]+nums[i+3]>target，说明此时剩下的三个数无论取什么值，四数之和一定大于 target，因此退出第一重循环；
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		// 在确定第一个数之后，如果 nums[i]+nums[n−3]+nums[n−2]+nums[n−1]<target，说明此时剩下的三个数无论取什么值，四数之和一定小于 target，因此第一重循环直接进入下一轮，枚举 nums[i+1]；
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		// 在确定前两个数之后，如果 nums[i]+nums[j]+nums[j+1]+nums[j+2]>target，说明此时剩下的两个数无论取什么值，四数之和一定大于 target，因此退出第二重循环；
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			// 在确定前两个数之后，如果 nums[i]+nums[j]+nums[n−2]+nums[n−1]<target，说明此时剩下的两个数无论取什么值，四数之和一定小于 target，因此第二重循环直接进入下一轮，枚举 nums[j+1]。
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}

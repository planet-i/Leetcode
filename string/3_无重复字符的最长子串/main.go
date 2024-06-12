/* 3. 无重复字符的最长子串
题目：给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
提示：
1. 0 <= s.length <= 5 * 104
2. s 由英文字母、数字、符号和空格组成。
*/

package main

// 左右指针滑动窗口法：
// 时间复杂度：O(N)，其中 N 是字符串的长度。左指针和右指针分别会遍历整个字符串一次。
// 空间复杂度：O(∣Σ∣)，其中 Σ 表示字符集（即字符串中可以出现的字符），∣Σ∣ 表示字符集的大小。
func lengthOfLongestSubstring(str string) int {
	n := len(str)
	window := make(map[byte]int, n) // 存窗口区间内，每个元素的数量
	left, right, res := 0, 0, 0
	for right < n {
		end := str[right]
		// 窗口扩大
		right++
		window[end]++
		// window[end] > 1就代表遇到重复字符了
		for window[end] > 1 { // 用for，是让左侧窗口收缩到这个被重复到的字符为止
			start := str[left]
			// 窗口缩小
			left++
			window[start]--
		}
		// 要求最长子串，得出最大的结果
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

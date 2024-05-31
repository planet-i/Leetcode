/* 3. 无重复字符的最长子串
题目：给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
提示：
1. 0 <= s.length <= 5 * 104
2. s 由英文字母、数字、符号和空格组成
----- 示例 1:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
----- 示例 2:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

package main

import (
	"fmt"
)

func main() {
	s := "abcbac"
	n := lengthOfLongestSubstring(s)
	fmt.Println(n)
}

func lengthOfLongestSubstring(str string) int {
	lens := len(str)
	window := make(map[byte]int, lens)
	left, right, res := 0, 0, 0
	for right < lens {
		b := str[right]
		window[b]++
		right++
		// left右滑动：window[b] > 1就代表遇到重复字符了
		for window[b] > 1 {
			a := str[left]
			window[a]--
			left++
		}
		// 判断区间是否满足条件
		if right-left > res {
			res = right - left
		}
	}
	return res
}

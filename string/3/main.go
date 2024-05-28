/*
3. 题目：给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
提示：
1. 0 <= s.length <= 5 * 104
2. s 由英文字母、数字、符号和空格组成
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

func lengthOfLongestSubstring(s string) int {
	lens := len(s)
	window := make(map[byte]int, lens)
	left, right, res := 0, 0, 0
	for right < lens {
		b := s[right]
		window[b]++
		right++
		// left右滑动：window[b] > 1就代表遇到重复字符了
		for window[b] > 1 {
			a := s[left]
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

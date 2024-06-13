/*
 76. 最小覆盖子串

题目：给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
*/
package main

import "math"

func minWindow(s string, t string) string {
	window := make(map[byte]int, len(s))
	need := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	ready, length := 0, math.MaxInt32 // 要找长度最小，就要先给他设置一个最大值，再等着被替换
	for right < len(s) {
		end := s[right]
		right++
		if _, ok := need[end]; ok {
			window[end]++
			// 开始我想的是根据题意要>=,但是当=的时候就已经符合条件了，如果>=,就会造成valid的异常增大
			if window[end] == need[end] {
				valid++
			}
		}
		for valid == len(need) { // 当窗口中已经包含了所有需要的字符时
			if right-left < length { // 更新最小覆盖子串长度及起始位置
				ready = left
				length = right - left
			}
			start := s[left]
			left++
			if _, ok := need[start]; ok {
				if window[start] == need[start] {
					valid--
				}
				window[start]--
			}
		}
	}
	if length == math.MaxInt32 { // 如果没有符合要求的子串，返回空字符串
		return ""
	}
	return s[ready : ready+length]
}

/*
 438. 找到字符串中所有字母异位词

题目：给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

注意：异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
提示:
- 1 <= s.length, p.length <= 3 * 104
- s 和 p 仅包含小写字母
*/

package main

// 左右指针滑动窗口法
func findAnagrams(s string, p string) []int {
	window := make(map[byte]int, len(s))
	need := make(map[byte]int, len(p))
	var res []int
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0 // valid代表有一个字符满足大小了
	for right < len(s) {
		end := s[right]
		right++
		if _, ok := need[end]; ok {
			window[end]++
			if window[end] == need[end] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
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
	return res
}

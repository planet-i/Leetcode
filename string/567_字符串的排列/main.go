/*
 567. 字符串的排列

题目：给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。

提示:
- 1 <= s.length, p.length <= 3 * 104
- s 和 p 仅包含小写字母
*/

package main

// 左右指针滑动窗口法
func checkInclusion(s1 string, s2 string) bool {
	window := make(map[byte]int, len(s2))
	need := make(map[byte]int, len(s1))
	for i := 0; i < len(s1); i++ {
		need[s1[i]] += 1
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		// 窗口右滑扩大
		end := s2[right]
		right++
		// 看这个值是否符合要求
		if _, ok := need[end]; ok {
			// 符合就将元素加入窗口
			window[end]++
			// 看元素个数是否符合要求
			if window[end] == need[end] {
				// 符合就代表找到一个了
				valid++
			}
		}
		// 当窗口扩大得要超过要求时，要进行判断了
		for right-left >= len(s1) {
			// 判断是否已经找到了符合条件的子串
			if valid == len(need) {
				return true
			}
			start := s2[left]
			left++
			// 看这个值是否符合要求
			if _, ok := need[start]; ok {
				// 看元素个数是否符合要求
				if window[start] == need[start] {
					// 符合就代表找到一个了
					valid--
				}
				// 符合就将元素移出窗口
				window[start]--
			}
		}
	}
	return false
}

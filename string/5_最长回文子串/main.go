/*
5. 最长回文子串
题目：给你一个字符串 s，找到 s 中最长回文子串。
*/
package main

func longestPalindrome(s string) string {
	res := ""
	// 每一轮得到的res是以s[i]为中心的最长回文串长度
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串。回文串长度是奇数。左右指针指向同一个唯一的中心
		s1 := palindrome(s, i, i)
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串。回文串长度是偶数。左右指针分别指向两个中心
		s2 := palindrome(s, i, i+1)
		// res = longest(res, s1, s2)，选出最长子串
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func palindrome(s string, left, right int) string {
	// 防止索引越界
	for left >= 0 && right < len(s) && s[left] == s[right] {
		// 向两边展开
		left--
		right++
	}
	// 返回以 s[left] 和 s[right] 为中心的最长回文串
	return s[left+1 : right]
}

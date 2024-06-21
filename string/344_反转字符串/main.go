/*
 344. 反转字符串

题目：编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
注意: 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
*/
package main

// 双指针法
// 时间复杂度：O(N)，其中 N 为字符数组的长度。一共执行了 N/2 次的交换。
// 空间复杂度：O(1)。只使用了常数空间来存放若干变量。
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

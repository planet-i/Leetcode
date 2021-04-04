// 242.有效的字母异位词：
// 给定两个字符串s和t，编写一个函数来判断t是否是s的字母异位词.假设字符串只包含小写字母
// 进阶：如果输入字符串包含Unicode字符怎么办?

/* 哈希映射：
首先判断两个字符串的长度是否相等，不相等返回false
相等：定义一个map
s负责将出现的字符value加加；t负责将出现的字符value减减
遍历map查看每个key是否为0，都为0则返回true */

package main

func main() {
	flag := isAnagram("中国 angle", "国中 ageln")
	println(flag)
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[byte]int) //s是作为ascii码存的,s[i]取出后是byte类型
	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}
	for _, elem := range count {
		if elem != 0 {
			return false
		}
	} //elem 乱序取出count中的值
	return true
}

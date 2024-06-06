/*
	189 轮转数组

题目: 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

	如果有两个中间结点，则返回第二个中间结点。

示例: 输入: nums = [1,2,3,4,5,6,7], k = 3

	输出: [5,6,7,1,2,3,4]
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 获取命令行参数
	args := os.Args

	// 如果没有足够的参数提供，提示用户正确的使用方法
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <k>")
		return
	}
	k, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error parsing k:", err)
		return
	}

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("法0", rotate(nums, k))
	fmt.Println("法1", rotate1(nums, k))
	rotate1_1(nums, k)
	fmt.Println("被改变后的切片", nums)
	fmt.Println("法2", rotate2(nums, k))
}

// 自己想的办法（不改变原数组）
func rotate(nums []int, k int) []int {
	n := len(nums)
	if k <= 0 || k%n == 0 || n <= 1 {
		return nums
	}
	p := k % n
	result := make([]int, n)
	// 存储后 p 位数据
	rotated := nums[n-p:]
	// 遍历切片，从第一位开始将当前位置的值，代替 i+p 位的值
	for i := n - 1; i >= p; i-- {
		result[i] = nums[i-p]
	}
	// 将 p 位数据放在切片的前几位
	for i := 0; i < p; i++ {
		result[i] = rotated[i]
	}
	return result
}

// rotate1_1 不返回，直接在原始切片上操作（改变原数组）
func rotate1_1(nums []int, k int) {
	n := len(nums)
	if k <= 0 || k%n == 0 || n <= 1 {
		return
	}
	p := k % n
	rotated := make([]int, n)
	// 将后面的元素移动到新切片的前面
	copy(rotated[p:], nums[:n-p])
	// 将前面的元素移动到新切片的后面
	copy(rotated[:p], nums[n-p:])
	// 将旋转后的切片拷贝回原始切片
	copy(nums, rotated)
}

// 法1：数组拷贝方式。找规律得出 result[(i + k) % nums.length] = nums[i]（不改变原数组）
func rotate1(nums []int, k int) []int {
	n := len(nums)
	if k <= 0 || k%n == 0 || n <= 1 {
		return nums
	}
	p := k % n
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[(i+p)%n] = nums[i]
	}
	return result
}

// 法2：数组翻转方式。将nums翻转，然后将数组分为两份。0~k-1 k~n-1 （改变原数组）
func rotate2(nums []int, k int) []int {
	n := len(nums)
	if k <= 0 || k%n == 0 || n <= 1 {
		return nums
	}
	// 计算旋转点
	p := k % n
	// 全量翻转
	reverse(nums)
	// 分别部分翻转
	reverse(nums[:p])
	reverse(nums[p:])
	return nums
}

// reverse 翻转数组
func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

// 方法三
func reverse3(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

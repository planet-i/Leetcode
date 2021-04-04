package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	a := fib1(n)
	b := fib2(n)
	c := fib3(n)
	fmt.Println(a, b, c)
}
func fib1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}
func fib2(n int) int { //自顶向下  从f(n)推出f(1),f(2) 递归
	if n < 1 {
		return 0
	}
	memo := make([]int, n+1)
	return helper(memo, n) //去从备忘录中取子问题值，解决重叠子问题
}
func helper(memo []int, n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = helper(memo, n-1) + helper(memo, n-2)
	return memo[n]
}

func fib3(n int) int { //自底向上 从f(1),f(2)推出f(n) 循环迭代
	if n < 1 {
		return 0
	}
	if n == 2 || n == 1 {
		return 1
	}
	prev, curr := 1, 1
	var sum int
	for i := 3; i <= n; i++ {
		sum = prev + curr
		prev = curr //只要存储n之前的两个状态即可
		curr = sum
	}
	return curr
}

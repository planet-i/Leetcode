package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	n := coinChange(coins, amount)
	fmt.Println(n)
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) //从1到amount+1，dp[i]就是amount=i的答案
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 //最大值是amount，所以答案都初始化为amount+1
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] > amount { //答案大于amount，则证明值没有被修改过，是无法被凑成的总金额
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

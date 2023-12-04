package main

import "fmt"

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	fmt.Println("法1", sumRange1(nums, 3, 5))

	prepare2(nums)
	fmt.Println("法2", sumRange2(3, 5))

	prepare3(nums)
	fmt.Println("法3", sumRange3(3, 5))
}

// 法1： 平均需要 O(n)
func sumRange1(nums []int, i int, j int) int {
	sum := 0
	for k := i; k <= j; k++ {
		sum += nums[k]
	}
	return sum
}

// 法2：为了避免多次调用使用相同的(i,j)造成重复计算，可以用二维数组存储所有(i,j)对应的值，调用时直接取。
var resArray2 [][]int

func prepare2(nums []int) {
	n := len(nums)
	res := make([][]int, n) // !!! len也要先给，不然无法直接赋值，会需要append进去
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
		sum := 0
		for j := i; j < n; j++ { // 需要 O(n2`)
			sum += nums[j]
			res[i][j] = sum
		}
	}
	resArray2 = res
	return
}

func sumRange2(i int, j int) int {
	return resArray2[i][j] // 需要 O(1)
}

// 法3：优化空间占用，二维数组变一维数组，优化时间复杂度。当前位置存前面位置的和 [0,i) 左开右闭 那就 0 ~ i-1
var resArray3 []int

func prepare3(nums []int) {
	n := len(nums)
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + nums[i] // 需要 O(n)
	}
	resArray3 = res
	return
}

func sumRange3(i int, j int) int {
	return resArray3[j+1] - resArray3[i] // 需要 O(1)
}

type NumArray struct {
	data []int
}

var res NumArray

func Constructor(nums []int) NumArray {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	res.data = preSum
	return res
}

func SumRange(left int, right int) int {
	return res.data[right+1] - res.data[left]
}

// 算法步骤
// 选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1； 即t越来越小，直到为1
// 按增量序列个数 k，对序列进行 k 趟排序；
// 每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。
// 仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。

package main

import (
	"fmt"
	"math"
)

type uint64Slice []uint64

func main() {
	numbers := []uint64{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}
	shellSort(numbers)
	fmt.Println(numbers)
}
func shellSort(numbers uint64Slice) {
	gap := 1
	for gap < len(numbers) {
		gap = gap*3 + 1
	} //不明白这里的gap是从哪里开始的
	for gap > 0 {
		for i := gap; i < len(numbers); i++ {
			tmp := numbers[i]
			j := i - gap //i和j相隔gap
			for j >= 0 && numbers[j] > tmp {
				numbers[j+gap] = numbers[j]
				j -= gap
			}
			numbers[j+gap] = tmp
		} //插入排序是将某个与前面的被选出来的某个位置换
		//希尔排序是与相差gap前的那个位置比较交换
		gap = int(math.Floor(float64(gap / 3))) //返回最大整数值
	}
}

//后使用的增量排序不会影响之前使用的增量排序

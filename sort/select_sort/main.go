// 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
// 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
// 重复第二步，直到所有元素均排序完毕。
package main

import "fmt"

type uint64Slice []uint64

func main() {
	numbers := []uint64{5, 23, 1, 6, 7, 9, 2}
	selectSort(numbers)
	fmt.Println(numbers)
}
func selectSort(numbers uint64Slice) {
	for i := 0; i < len(numbers)-1; i++ {
		min := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[min] {
				min = j //找到了最小值的位置
			}
		}
		if i != min {
			numbers.swap(i, min)
		} //确保最小值是从前往后排列
	}
}
func (numbers uint64Slice) swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}

// 算法步骤
// 从数列中挑出一个元素，称为 “基准”（pivot）;
// 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
// 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
package main

import "fmt"

type uint64Slice []uint64

func main() {
	numbers := []uint64{5, 4, 20, 3, 8, 2, 8}
	quickSort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers) //切片可以直接print打印？？
}

func quickSort(numbers uint64Slice, start, end int) {
	var middle int
	tempStart := start
	tempEnd := end
	if tempStart >= tempEnd {
		return
	}
	pivot := numbers[start] //确认一个基准数
	for start < end {
		for start < end && numbers[end] > pivot {
			end-- //从后往前找 比pivot小的
		}
		if start < end {
			numbers.swap(start, end)
			start++
		} // 将小的一道前面，基准换到end位置上
		for start < end && numbers[start] < pivot {
			start++ //从前往后找 比pivot大的
		}
		if start < end {
			numbers.swap(start, end)
			end--
		} //将大的移到后面，基准换到start位置上
	}
	//基准值一直没变，元素位置在交换，前后指针在交替着往中间靠拢
	// 基准值在end指针和start指针位置徘徊
	// 前后指针移动，将基准值与比较值不断交换位置
	numbers[start] = pivot // 为什么这里要赋值，他们已经重合啦，本来中间相遇的位置就会是基准值在这里
	// 如果这里本来不是基准值，而随意赋值的话，就会改变本来的数组元素组成呀？？
	middle = start // end和start在中间相遇

	quickSort(numbers, tempStart, middle-1) // 左序列递归
	quickSort(numbers, middle+1, tempEnd)   // 右序列递归
}
func (numbers uint64Slice) swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}

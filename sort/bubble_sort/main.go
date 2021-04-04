package main

import "fmt"

type uint64Slice []uint64

func main() {
	numbers := []uint64{5, 4, 2, 3, 8} // 为什么用uint64的切片
	sortBubble(numbers)
	fmt.Println(numbers)
}

func sortBubble(numbers uint64Slice) {
	length := len(numbers)
	if length == 0 {
		return
	}
	flag := true //
	for i := 0; i < length-1 && flag; i++ {
		flag = false
		for j := 0; j < length-1-i; j++ {
			if numbers[j] > numbers[j+1] { //大的去后面
				numbers.swap(j, j+1)
				flag = true // 有交换 flag为true， 无交换则证明此轮已经有序，flag为flase 跳出循环
			}
		}
	}
}
func (numbers uint64Slice) swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}

package main

import "fmt"

func main() {
	var data = []int{38, 65, 97, 76, 13, 27, 49}
	insertSort(data)
}

func insertSort(data []int) []int {
	n := len(data)
	// 边界检查
	if n < 2 {
		return data
	}
	// 遍历无序列表，每次取第一个值做key去插入有序列表  1～n-1（忽略头，一个肯定有序）
	for i := 1; i < n; i++ {
		key := data[i]
		// data[i] 与 “data[i - 1]...data[0]“ 比较
		j := i - 1
		for j >= 0 && data[j] > key {
			// 有序列表中比插入值大的后移
			data[j+1] = data[j]
			j--
		}
		// j+1 就是该插入key的位置
		data[j+1] = key
		fmt.Println(data)
	}

	return data
}

// var j int
// for j = i - 1; j >= 0; j-- {
// 	// 如果当前值小于前一个值
// 	if data[j] > key {
// 		data[j+1] = data[j]
// 	} else {
// 		// 证明已经插入到合适的位置直接跳出即可
// 		break
// 	}
// }

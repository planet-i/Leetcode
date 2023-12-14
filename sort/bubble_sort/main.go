package main

func main() {
	var data = []int{38, 65, 97, 76, 13, 27, 49}
	bubbleSort(data)
}

func bubbleSort(data []int) []int {
	n := len(data)
	if n < 2 {
		return data
	}
	// n-1趟
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

func bubbleSort2(data []int) []int {
	n := len(data)
	if n < 2 {
		return data
	}
	// 用来标记是否中间轮次就已经有序了
	flag := true
	for i := 0; i < n-1 && flag; i++ {
		flag = false
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] { // 大的去后面
				data[j], data[j+1] = data[j+1], data[j]
				flag = true // 有交换 flag为true，无交换则证明此轮已经有序，flag为flase 跳出循环
			}
		}
	}
	return data
}

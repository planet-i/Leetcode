package main

import "fmt"

func main() {
	var data = []int{3, 7, 9, 4, 10}
	heapSort(data)
}

// 小顶堆-降序、大顶堆-升序
func heapSort(data []int) {
	n := len(data)
	// 构建堆：从最后一个非叶子结点开始向前遍历。（从前往后会破坏前面已经调整好的子堆）
	fmt.Println("初始值 ", data)
	for i := n/2 - 1; i >= 0; i-- {
		fmt.Println("构建中-", data[i])
		minHeap(data, i, n-1)
	}
	fmt.Println("构建后 ", data)
	// 取出元素，调整堆
	for i := n - 1; i >= 0; i-- {
		fmt.Println("前-->", data)
		data[0], data[i] = data[i], data[0] // 将堆顶元素与数组的末尾元素交换
		fmt.Println("后---", data)
		minHeap(data, 0, i-1) // 缩小需要调整的范围。根都是从0开始，限制长度为i. i是从数组的后往前一个个缩小
	}
}

// 初始值 [3 7 9 4 10]
// 构建后 [3 4 9 7 10]
// 前--> [3 4 9 7 10]
// 后--- [10 4 9 7 3]
// 前--> [4 7 9 10] 3
// 后--- [10 7 9 4] 3
// 前--> [7 10 9] 4 3
// 后--- [9 10 7] 4 3
// 前--> [9 10] 7 4 3
// 后--- [10 9] 7 4 3
// 前--> [10] 9 7 4 3
// 后--- [10] 9 7 4 3

// 构建小顶堆
func minHeap(data []int, pos, n int) {
	var temp int
	child := 0
	for temp = data[pos]; 2*pos+1 <= n; pos = child {
		child = 2*pos + 1
		if child < n && data[child] > data[child+1] {
			child++
		}
		if data[child] < temp {
			data[pos] = data[child]
		} else {
			break
		}
	}
	data[pos] = temp
}

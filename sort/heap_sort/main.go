package main

import "fmt"

func main() {
	var data = []int{3, 7, 9, 4, 10}
	heapSort(data)
	fmt.Println("-----===========")
	var data2 = []int{38, 65, 97, 76, 13, 27, 49}
	heapSort2(data2, 0, len(data2))
}

// 小顶堆-降序
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

// 大顶堆-升序
func heapSort2(data []int, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// 构建堆
	fmt.Println("初始值 ", data)
	for i := (hi - 1 - 1) / 2; i >= 0; i-- {
		fmt.Println("构建中-", i)
		siftDown(data, i, hi, first)
	}

	// 取出元素，调整堆
	for i := hi - 1; i >= 0; i-- {
		fmt.Println("前-->", data)
		data[first], data[first+i] = data[first+i], data[first] // 将堆顶元素与数组的末尾元素交换
		fmt.Println("后---", data)
		siftDown(data, lo, i, first) // 缩小需要调整的范围。根都是从0开始，限制长度为i. i是从数组的后往前一个个缩小
	}
}

func siftDown(data []int, lo, hi, first int) {
	// 初始化根节点,表示当前子树的根节点
	root := lo
	// 循环向下调整,直到满足堆的性质。
	for {
		child := 2*root + 1
		// 判断当前节点的左子节点是否超出了实际数据范围，超出可能会访问到不存在的元素，造成程序错误
		if child >= hi {
			break
		}
		// 如果存在右子节点，并且右子节点比左子节点大，则选择右子节点作为需要交换的节点。
		if child+1 < hi && (data[first+child] < data[first+child+1]) {
			child++
		}
		// 判断是否需要交换：比较根节点和子节点的大小，如果根节点比子节点大，则表示已经满足堆的性质，可以退出循环。
		if !(data[first+root] < data[first+child]) {
			return
		}
		// 交换节点：如果根节点比子节点小，则交换根节点和子节点的值，并更新 root 为子节点的索引。
		data[first+root], data[first+child] = data[first+child], data[first+root]
		root = child
	}
}

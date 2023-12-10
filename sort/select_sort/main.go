package main

func main() {
	var data = []int{38, 65, 97, 76, 13, 27, 49}
	selectSort(data)
	selectSort2(data)
}

func selectSort(data []int) []int {
	n := len(data)
	// 边界检查
	if n < 2 {
		return data
	}
	// 遍历无序列表，找到其中最小值，放到i位置。0～n-2（忽略尾，最后剩下的就是最大的）
	for i := 0; i < n-1; i++ {
		// 从第一个起，依次比较，获取实际最小值索引
		min := i
		for j := i + 1; j < n; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		// 将最小值放有序列表里
		if min != i {
			data[i], data[min] = data[min], data[i]
		}
	}
	return data
}

func selectSort2(data []int) []int {
	n := len(data)
	// 边界检查，直接返回
	if n < 2 {
		return data
	}
	left, right := 0, n-1
	for left < right {
		min := left
		max := right
		for i := left; i <= right; i++ {
			// 标记较小值和较大值
			if data[i] < data[min] {
				min = i
			}
			if data[i] > data[max] {
				max = i
			}
		}
		// 最小值放最左端
		data[left], data[min] = data[min], data[left]
		// 计算得出的最大值在left位置，但是上一步，原来left位置的值被放到min位置上。
		if max == left {
			max = min
		}
		// 最大值放最右端
		data[right], data[max] = data[max], data[right]
		// 缩小无序列表
		left++
		right--
	}
	return data
}

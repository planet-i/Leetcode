//初始时，a[0]自成1个有序区，无序区为a[1..n-1]。
//从头到尾依次扫描无序表，将这个元素插入到有序表中，使之成为新的有序表.
//查找插入的位置采用二分查找法，不断比较有序表的中间位置的值，找出待插入元素可能待在的区间，重复n-1次完成整个排序过程。
package main

import "fmt"

type uint64Slice []uint64

func main() {
	numbers := []uint64{5, 4, 20, 3, 8, 2, 9}
	insertSort(numbers)
	fmt.Println(numbers)
}

func insertSort(numbers uint64Slice) {
	for i := 1; i < len(numbers); i++ { //遍历无序表
		tmp := numbers[i] //从无序表中取出第一个数
		//从待排序序列开始比较,找到比其小的数
		j := i
		for j > 0 && tmp < numbers[j-1] {
			numbers[j] = numbers[j-1] //说明tmp应该放在j之前的位置上
			j--                       //下标向前移动，找到准确位置
		}
		// 存在比其小的数插入
		if j != i {
			numbers[j] = tmp // 把tmp放到j位置上
		}
	}
}

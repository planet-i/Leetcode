/* 21 合并两个有序链表
题目: 将两个升序链表合并为一个新的 升序 链表并返回。
	  新链表是通过拼接给定的两个链表的所有节点组成的。
重点: 使⽤虚拟头结点可以简化边界情况的处理。
	  双指针
*/
package main

import "fmt"

func main() {
	// 创建两个链表用于测试
	list1 := &ListNode{1, &ListNode{2, &ListNode{6, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	// 调用合并链表的函数
	mergedList := mergeTwoLists2(list1, list2)

	// 输出合并后的链表
	fmt.Println("Merged list:")
	for p := mergedList; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	// 一开始设置一个虚拟节点，它的值为 -1，它的值可以设置为任何的数，因为我们根本不需要使用它的值
	head := &ListNode{-1, nil}
	// 设置一个指针，指向虚拟节点
	pre := head
	// 通过一个循环，不断的比较 l1 和 l2 中当前节点值的大小，直到 l1 或者 l2 遍历完毕为止
	for list1 != nil && list2 != nil {
		// 让 pre 的 next 指针指向更小值的节点，并将对应链表的指针后移
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next // 将list1后移
		} else {
			pre.Next = list2
			list2 = list2.Next // 将list2后移
		}
		// 让 pre 向后移动
		pre = pre.Next
	}

	// 如果 list1 中还有节点，或者一开始 list2 就是nil
	if list1 != nil {
		pre.Next = list1
	}
	// 如果 list2 中还有节点，或者一开始 list1 就是nil
	if list2 != nil {
		pre.Next = list2
	}
	// 最后返回虚拟节点的 Next 指针
	return head.Next
}

/* 86 分隔链表
题目: 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，
	  使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
      你应当 保留 两个分区中每个节点的初始相对位置。
重点: 一边小一边大，比较大小后分解，两个链表最后合起来
	  保持相对位置，则只要简单从头到尾遍历，不需要分别有序再合并
*/
package main

import "fmt"

func main() {
	// 创建测试链表
	list := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}

	// 调用分隔链表的函数
	resList := partition(list, 3)

	for p := resList; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{-1, nil}
	dummy2 := &ListNode{-1, nil}
	p1 := dummy1
	p2 := dummy2
	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		// 断开原链表中的每个节点的 next 指针
		temp := head.Next
		head.Next = nil
		head = temp
	}
	// 连接两个链表
	p1.Next = dummy2.Next

	return dummy1.Next
}

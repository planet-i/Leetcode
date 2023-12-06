/* 876 链表的中间结点
题目: 给你单链表的头结点 head ，请你找出并返回链表的中间结点。
	  如果有两个中间结点，则返回第二个中间结点。
重点: 返回不是单个值，而是以这个结点开头的链表。快慢指针，快的走两步，慢的走一步
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 创建链表用于测试
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	// 调用函数
	res := middleNode(list)

	// 输出
	for p := res; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		// 慢指针走一步，快指针走两步
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 慢指针指向中点
	return slow
}

/* 19 删除链表的倒数第 N 个结点
题目: 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
重点: 使⽤虚拟头结点可以简化边界情况的处理。因为第一个结点前面没有结点了
	  快慢指针相差 k 步，快指针到最后，慢指针到倒数 k 个的位置。要删除结点，要找到被删除结点的前一个结点
*/
package main

import "fmt"

func main() {
	// 创建链表用于测试
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	// 调用函数
	res := removeNthFromEnd(list, 2)

	// 输出
	for p := res; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 构建虚拟头结点，如果要删除的是第一个结点，就不再需要特别判断
	dummy := &ListNode{-1, head}
	// 删除倒数第 n 个，要先找倒数第 n + 1 个结点
	x := findFromEnd(dummy, n+1)
	// 删掉倒数第 n 个结点
	x.Next = x.Next.Next
	return dummy.Next
}

func findFromEnd(head *ListNode, k int) *ListNode {
	p1 := head
	// p1 先走 k 步
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	p2 := head
	// p1 和 p2 同时走 n - k 步
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	// p2 现在指向第 n - k + 1 个节点，即倒数第 k 个节点
	return p2
}

package main

import "fmt"

func main() {
	//初始化一个链表用于测试
	head := &ListNode{Val: -1}
	head.Next = &ListNode{Val: 5}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 0}
	//输出链表
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, "\t")
		cur = cur.Next
	}
	fmt.Println()
	//链表排序
	head = sortList(head)
	//输出链表
	cur = head
	for cur != nil {
		fmt.Print(cur.Val, "\t")
		cur = cur.Next
	}
	fmt.Println()
}

//定义链表的头节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	ret, _ := process(head)
	return ret
}
func process(head *ListNode) (*ListNode, *ListNode) {
	left, leftend, mid, midend, right, rightend := partition(head)
	if left != nil {
		left, leftend = process(left)
	}
	if right != nil {
		right, rightend = process(right)
	}
	return merge(left, leftend, mid, midend, right, rightend)
}
func partition(head *ListNode) (*ListNode, *ListNode, *ListNode, *ListNode, *ListNode, *ListNode) {
	left := &ListNode{} //虚拟节点
	leftend := left
	mid := head
	midend := mid
	right := &ListNode{} //虚拟节点
	rightend := right

	cur := head.Next
	for cur != nil {
		if cur.Val < mid.Val {
			leftend.Next = cur
			leftend = leftend.Next
		} else if cur.Val == mid.Val {
			midend.Next = cur
			midend = midend.Next
		} else {
			rightend.Next = cur
			rightend = rightend.Next
		}
		cur = cur.Next
	}
	leftend.Next = nil
	midend.Next = nil
	rightend.Next = nil

	left = left.Next
	if left == nil {
		leftend = nil
	}
	right = right.Next
	if right == nil {
		rightend = nil
	}
	return left, leftend, mid, midend, right, rightend
}
func merge(left, leftend, mid, midend, right, rightend *ListNode) (*ListNode, *ListNode) {
	head := &ListNode{}
	headend := head

	if left != nil {
		headend.Next = left
		headend = leftend
	}
	headend.Next = mid
	headend = midend
	if right != nil {
		headend.Next = right
		headend = rightend
	}
	head = head.Next
	if head == nil {
		headend = nil
	}
	return head, headend
}

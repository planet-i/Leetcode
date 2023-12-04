package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	printlnLinkNodeList(head)
	k := 6
	fmt.Println("k =", k)
	fmt.Println("----------------")
	ret := rotateRight(head, k)
	printlnLinkNodeList(ret)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	//找尾节点并且计数
	cnt := 1
	tail := head
	for tail.Next != nil {
		cnt++
		tail = tail.Next
	}

	k = k % cnt
	if k == 0 { //刚好是头节点，就不用操作了。
		return head
	}

	//找倒数第k+1节点
	fast := head
	slow := head
	k++
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	//缓存结果
	ans := slow.Next
	//尾节点连头节点
	tail.Next = head
	//倒数k+1节点无Next指针
	slow.Next = nil

	return ans
}

//链表打印
func printlnLinkNodeList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, "\t")
		cur = cur.Next
	}
	fmt.Println()
}

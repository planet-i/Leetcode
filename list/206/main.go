package main

import (
	"fmt"
)

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{val: arr[0]}
	current := head

	for i := 1; i < len(arr); i++ {
		current.next = &ListNode{val: arr[i]}
		current = current.next
	}

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.val)
		if head.next != nil {
			fmt.Print(" -> ")
		}
		head = head.next
	}
	fmt.Println()
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	head := createList(arr)
	fmt.Print("Original List: ")
	printList(head)

	reversedHead := reverseList(head)
	fmt.Print("Reversed List: ")
	printList(reversedHead)
	test()
}

type ListNode struct {
	val  int
	next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead *ListNode

	for head != nil {
		// node := head.next   // 1
		// head.next = newHead // 2     实现反转
		// newHead = head      // 3     实现第一个结点往后移
		// head = node         // 4     1和4实现第二个结点往后移
		head, head.next, newHead = head.next, newHead, head
	}
	return newHead
}

func test() {
	a := "a"
	b := "b"
	c := "c"
	//a, b, c = b, c, a
	tmp := b
	b = c
	c = a
	a = tmp
	fmt.Println(a, b, c)
}

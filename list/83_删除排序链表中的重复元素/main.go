/*
 83. 删除排序链表中的重复元素

题目：给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// len(nums) == 0
	if head == nil {
		return nil
	}
	// slow, fast := 0, 0
	slow, fast := head, head
	// fast < len(nums)
	for fast != nil {
		if slow.Val != fast.Val {
			// nums[slow] = nums[fast];
			slow.Next = fast
			// slow++;
			slow = slow.Next
		}
		// fast++
		fast = fast.Next
	}
	// 断开与后面重复元素的连接
	slow.Next = nil
	return head
}

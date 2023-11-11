/* 160 相交链表
题目描述：两个单链表的头节点 headA和 headB，请找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null,其中链式结构中不存在环（像个y的形状）
解题链接：https://mp.weixin.qq.com/s/IZSwAqY_I0KQfsdtowQO3Q
解题要点：a+(b-c) = b+(a-c)
1. 双指针法: 分别遍历链表 A 和链表 B，当其中一个链表遍历结束后，将指针重新指向另一个链表的头部，继续遍历。
            两个指针分别走过的路径长度为 a+b 和 b+a，它们会在相交节点相遇，或者同时到达链表尾部（此时没有相交节点）。
			时间复杂度为 O( a + b )。空间复杂度为 O(1)。
2. 哈希法: 链表的节点只有一个指向，则如果两个链表存在相同的节点则为它们的相交节点
*/
package main

import "fmt"

func main() {
	testTwoList()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode 两个单链表相交的起始节点（双指针法）
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}
		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}
	return pa
}

// getIntersectionNodeForHash 两个单链表相交的起始节点（哈希法）
func getIntersectionNodeForHash(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

func testTwoList() {
	// 创建链表A
	node1 := &ListNode{Val: 2, Next: nil}
	node2 := &ListNode{Val: 6, Next: nil}
	node3 := &ListNode{Val: 4, Next: nil}
	node1.Next = node2
	node2.Next = node3
	// 创建链表B
	node4 := &ListNode{Val: 1, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	node4.Next = node5
	// 使链表B与链表A相交于节点3    2 6  4
	node5.Next = node3 //       1 5 /
	// 两个链表相交节点的值
	intersectionNode := getIntersectionNode(node1, node4)
	if intersectionNode != nil {
		fmt.Println("相交节点的值为:", intersectionNode.Val)
	} else {
		fmt.Println("链表A与链表B不相交")
	}
	intersectionNodeForHash := getIntersectionNodeForHash(node1, node4)
	if intersectionNodeForHash != nil {
		fmt.Println("相交节点的值为:", intersectionNode.Val)
	} else {
		fmt.Println("链表A与链表B不相交")
	}
	return
}

package linkedlist

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/problems/linked-list-cycle/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func Test_hasCycle(t *testing.T) {
	var node1, node2 *ListNode

	node2 = &ListNode{Val: 2}
	node1 = &ListNode{Val: 1}
	node2.Next = node1
	node1.Next = node2

	fmt.Println(node1, node2)

	fmt.Println("is it has cycle:", hasCycle(node1))
}

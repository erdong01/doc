package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {

	var head ListNode
	head = ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	fmt.Println("head", head)
	newHead := reverseList(&head)
	fmt.Println(newHead)
}
func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

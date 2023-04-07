package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var list1 *ListNode = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	var list2 *ListNode = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	// var list1 *ListNode = &ListNode{
	// 	Val:  2,
	// 	Next: nil,
	// }

	// var list2 *ListNode = &ListNode{
	// 	Val:  1,
	// 	Next: nil,
	// }
	res := mergeTwoLists(list1, list2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head = &ListNode{}
	var tail *ListNode = head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next, tail, list1 = list1, list1, list1.Next
		} else {
			tail.Next, tail, list2 = list2, list2, list2.Next
		}
	}
	if list2 != nil {
		tail.Next = list2
	}
	if list1 != nil {
		tail.Next = list1
	}
	return head.Next
}

func pNode(res *ListNode) {
	for res != nil {
		fmt.Println("val:", res.Val)
		res = res.Next
	}
}

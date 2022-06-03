package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var treeNode TreeNode
	treeNode.Val = 3

	TreeNode1 := TreeNode{
		Val: 9,
	}
	treeNode.Left = &TreeNode1

	TreeNode3 := TreeNode{
		Val: 15,
	}
	TreeNode4 := TreeNode{
		Val: 7,
	}
	TreeNode2 := TreeNode{
		Val:   20,
		Left:  &TreeNode3,
		Right: &TreeNode4,
	}
	treeNode.Right = &TreeNode2
	maxDepth(&treeNode)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lhight := maxDepth(root.Left)
	rhight := maxDepth(root.Right)
	fmt.Println(lhight, rhight)
	return max(lhight, rhight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

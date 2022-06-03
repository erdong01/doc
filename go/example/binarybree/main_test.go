package binarybree

import (
	"fmt"
	"testing"
)

//二叉树
func TestXxx(t *testing.T) {
	root := &Hero{
		No:   1,
		Name: "松江",
	}
	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}
	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}
	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}
	right1.Right = right2
	root.Left = left1

	node10 := &Hero{
		No:   10,
		Name: "林冲",
	}
	node12 := &Hero{
		No:   12,
		Name: "林冲",
	}
	left1.Left = node10
	left1.Right = node12
	root.Right = right1

	PreOrder(root)
}

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no %d name %s \n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no %d name %s \n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}
func PostOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		InfixOrder(node.Right)
		fmt.Printf("no %d name %s \n", node.No, node.Name)
	}
}

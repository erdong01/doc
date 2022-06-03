package main

import (
	"fmt"
)

func main() {
	var head *CatNode = &CatNode{}
	InsertCatNode(head, &CatNode{no: 1, name: "猫1"})
	InsertCatNode(head, &CatNode{no: 2, name: "猫2"})
	InsertCatNode(head, &CatNode{no: 3, name: "猫3"})
	ListCircleLink(head)

	DelCatNode(head, 2)
	ListCircleLink(head)
}

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		return
	}

	temp := head
	for {
		if temp.next == head {
			temp.next = newCatNode
			newCatNode.next = head
			break
		}
		temp = temp.next
	}
}

func ListCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空的环形列表")
		return
	}

	for {
		fmt.Printf("猫信息：%d %s \n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func DelCatNode(head *CatNode, id int) {

	if head.next == nil {
		return
	}
	if head.next == head {
		if head.no != id {
			return
		}
		head.no = 0
		head.name = ""
		head.next = nil
		return
	}

	temp := head
	helper := head
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	for {
		if temp.next == head {
			if temp.no == id {
				helper.next = temp.next
			}
			break
		}
		if temp.no == id {
			if temp == head {
				head.no = head.next.no
				head.name = head.next.name
				head.next = head.next.next
				break
			}
			helper.next = temp.next
			break
		}
		temp = temp.next
		helper = helper.next
	}

}

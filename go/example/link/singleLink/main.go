package main

import "fmt"

func main() {
	head := &HeroNode{}
	InsertHeroNode(head, &HeroNode{no: 1, name: "宋江", nickName: "及时雨"})
	InsertHeroNode(head, &HeroNode{no: 2, name: "林冲", nickName: "禁军教头"})
	InsertHeroNode(head, &HeroNode{no: 3, name: "吴智", nickName: "智多星"})
	ListHeroNode(head)

}

type HeroNode struct {
	no       int
	name     string
	nickName string
	next     *HeroNode
	tail     *HeroNode
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	if head.tail == nil {
		head.next = newHeroNode
		head.tail = newHeroNode
	} else {
		head.tail.next = newHeroNode
		head.tail = newHeroNode
	}
}

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if flag == false {
		fmt.Println("英雄已存在！")
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

//显示所有链表节点信息

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空列表")
		return
	}
	for {
		fmt.Printf("[%d , %s ,%s ,%p ]", temp.next.no, temp.next.name, temp.next.nickName, temp.next)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

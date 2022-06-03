package main

import (
	"fmt"
)

//双向链表
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
	pre      *HeroNode
	next     *HeroNode
	tail     *HeroNode
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	if head.tail == nil {
		newHeroNode.pre = head
		head.next = newHeroNode
		head.tail = newHeroNode
	} else {
		head.tail.next = newHeroNode
		newHeroNode.pre = head.tail
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
	if !flag {
		fmt.Println("英雄已存在！")
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}

}

func DelHeroNode(head *HeroNode, no int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.no == no {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("sorry,要删除的id不存在")
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
func ListHeroNode2(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空列表")
		return
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	for {
		fmt.Printf("[%d , %s ,%s ,%p ]", temp.no, temp.name, temp.nickName, temp)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

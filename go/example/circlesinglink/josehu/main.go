package main

import "fmt"

func main() {
	first := AddBoy(5)
	ShowBoy(first)
	PlayGame(first, 2, 3)
}

type Boy struct {
	No   int
	Next *Boy
}

func AddBoy(num int) *Boy {
	if num < 1 {
		return nil
	}

	var first *Boy = &Boy{}
	var curBoy *Boy = &Boy{}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

func ShowBoy(first *Boy) {
	if first.Next == nil {
		return
	}
	curBoy := first

	for {
		fmt.Printf("小孩编号=%d \n", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func PlayGame(first *Boy, startNo int, countNum int) {
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	for i := 1; i <= startNo-1; i++ {
		fmt.Println(i)
		first = first.Next
		tail = tail.Next
	}
	for {
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.No)
		first = first.Next
		tail.Next = first
		if tail == first {
			break
		}
	}
	fmt.Printf("最后小孩编号为%d 出圈 \n", first.No)
}

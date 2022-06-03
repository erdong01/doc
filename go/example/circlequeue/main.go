package main

import (
	"errors"
	"fmt"
)

func main() {
	circleQueue := &CircleQueue{
		maxSize: MaxSize,
		head:    0,
		tail:    0,
	}
	circleQueue.Push(1)
	circleQueue.Push(2)
	circleQueue.Push(3)
	circleQueue.Push(4)
	i, e := circleQueue.Pop()
	fmt.Println(i, e)
}

const MaxSize = 5

type CircleQueue struct {
	maxSize int
	array   [MaxSize]int
	head    int //指向队列队首
	tail    int //指向队列队尾
}

//入队列 addQueue

func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列已满")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize

	return
}

func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("队列已空")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

func (this *CircleQueue) ListQueue() {
	fmt.Println("队列打印：")
	//取出当前队列
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d] = %d \n", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println("打印结束")
}

//判断环形队列为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//判断环形队列是空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

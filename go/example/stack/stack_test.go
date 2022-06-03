package stack

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

type Stack struct {
	MaxTop int //表示最可以存放个数
	Top    int //表示栈顶
	arr    [20]int
}

func (this *Stack) Push(val int) (err error) {
	//先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满了")
		return errors.New("stack full")
	}
	this.Top++

	this.arr[this.Top] = val
	return nil
}

func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		return -1, errors.New("没有数据了")
	}
	val = this.arr[this.Top]
	this.Top--
	return
}

func (this *Stack) List() {
	if this.Top == -1 {
		return
	}
	var curTop = this.Top
	for i := curTop; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d \n", i, this.arr[i])
	}
}

func TestXxx(t *testing.T) {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.List()
	val, _ := stack.Pop()
	fmt.Println(val)
	stack.List()
}

func TestExpStack(t *testing.T) {
	numStack := &Stack{MaxTop: 20, Top: -1}
	operStack := &Stack{MaxTop: 20, Top: -1}
	ex := "3+2*6-2"

	index := 0

	num1 := 0
	num2 := 0
	oper := 0
	result := 0

	keepNum := ""
	for {
		ch := ex[index : index+1]
		temp := int([]byte(ch)[0])
		if operStack.IsOper(temp) {
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					numStack.Push(result)
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else {
			keepNum += ch
			if index == len(ex)-1 {

				ii, _ := strconv.Atoi(keepNum)
				numStack.Push(ii)
			} else {
				if operStack.IsOper(int([]byte(ex[index+1 : index+2])[0])) {
					ii, _ := strconv.Atoi(keepNum)
					numStack.Push(ii)
					keepNum = ""
				}
			}

		}
		if index == numStack.MaxTop {
			break
		}
		index++
	}
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		numStack.Push(result)
	}
	fmt.Println("result结果", result)
}

//判断是否是运算符

func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 { //乘法
		return true
	}
	return false
}

func (this *Stack) Cal(numl int, num2 int, oper int) int {
	res := 0
	switch res {
	case 42:
		res = num2 * numl
	case 43:
		res = num2 * numl
	case 45:
		res = num2 * numl
	case 47:
		res = num2 * numl
	}

	return res
}

func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 2
	}
	return res
}

func TestValid(t *testing.T) {
	v:=  13579.55 +13584.51 + 13050.75+13651.23+13878.18+13951.93+13914.29+3240
	fmt.Println(v)

}

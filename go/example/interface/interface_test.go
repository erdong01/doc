package interface_test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	c := Work{3}
	var a A = &c
	var b B = &c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}

type A interface {
	ShowA() int
}
type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w *Work) ShowA() int {
	return w.i + 10
}

func (w *Work) ShowB() int {
	return w.i + 20
}

func TestCalc(t *testing.T) {
	a:=1
	b:= 2
	defer calc("1",a,calc("10",a,b))
	a =0 
	defer calc("2",a ,calc("20",a,b))
	b=1
	//20 0 2  2
	//2 0  2  2  
	//10 1 2 3
	//1 1  3 4
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

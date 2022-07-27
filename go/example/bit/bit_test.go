package bit

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var a uint = 60
	var b uint = 13
	var c uint

	c = a & b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	fmt.Printf("第一行 - c 的值为 %d\n", b<<2)
}

const (
	a = 11
	b = 123
	c = 342
	d = 421
)

func TestConst(t *testing.T) {
	fmt.Printf("%b \n", a)
	fmt.Printf("%b \n", b)
	fmt.Printf("%b \n", c)
	fmt.Printf("%b \n", d)

	f := a | b | c | d
	fmt.Println("f , d", f, d)
	fmt.Println(f & d)
	f = f ^ d
	fmt.Println(f & d)
}

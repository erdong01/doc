package defer_test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	m := 10
	defer fmt.Printf("first defer %d \n", +m)
	m = 100
	defer func() {
		fmt.Printf("second defer %d \n", m)
	}()
	m *= 10
	defer fmt.Printf("third defer %d \n", m)
	funcVal := func1()
	funcVal()
	m *= 10
}

func func1() func() {
	defer fmt.Println("before return")
	return func() {
		defer fmt.Println("in the return")
	}
}

func TestF(t *testing.T) {
	sum := f(3)
	fmt.Println("F", sum)
	p := 1
	incr(&p)
	fmt.Println(p)
}

func f(n int) (r int) {
	defer func() {
		r += n

		recover()
	}()
	var f func()
	defer f()
	f = func() {
		r += 2
		fmt.Println(r)
	}
	return n + 1
}

func incr(p *int) int {
	*p++
	return *p
}

func TestDeferPanic(t *testing.T) {
	defer fmt.Println("打印前")
	defer fmt.Println("打印中")
	defer fmt.Println("打印后")
	panic("触发异常")
}

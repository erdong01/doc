package closure_test

import (
	"fmt"
	"testing"
)

func TestMain2(m *testing.T) {
	printFunc2()
}
func printFunc2() {
	fmt.Println(func2())
}
func func3() int {
	val := 10
	defer func() {
		val += 1
	}()
	return 100
}
func func2() (val int) {
	val = 10
	defer func() {
		val += 1
	}()
	return 100
}
func TestMain(m *testing.T) {
	printFunc1()
}

func printFunc1() {
	sumFunc := func1()
	fmt.Println(sumFunc(1))
	fmt.Println(sumFunc(2))
}

func func1() func(val int) int {
	sum := 0
	return func(val int) int {
		sum += val
		return sum
	}
}

func TestXxx111(t *testing.T) {
	funcs := testFunc(100)
	for _, v := range funcs {
		v()
	}
}

func testFunc(x int) []func() {
	var funcs []func()
	values := []int{1, 2, 3}
	for _, val := range values {
		funcs = append(funcs, func() {
			fmt.Printf("testFunc val = %d\n", x+val)
		})
	}
	return funcs
}

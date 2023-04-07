package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	fmt.Println(F(1, 1, 1))
	fmt.Println(F(2, 1, 1))
	fmt.Println(F(3, 1, 1))
	fmt.Println(F(4, 1, 1))
	fmt.Println(F(5, 1, 1))
}

func F(n int, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	return F(n-1, a2, a1+a2)
}

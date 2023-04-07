package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 1}
	ans := largestRectangleArea(data)
	fmt.Println(ans)

}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, rgiht := make([]int, n), make([]int, n)
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	mono_stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			rgiht[i] = n
		} else {
			rgiht[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	fmt.Println("rgiht", rgiht)
	fmt.Println("left", left)
	for i := 0; i < n; i++ {
		fmt.Println(ans, (rgiht[i]-left[i]-1)*heights[i], rgiht[i], left[i], heights[i])
		ans = max(ans, (rgiht[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

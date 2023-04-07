package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var nums = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(nums)
	fmt.Println(res)
}

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

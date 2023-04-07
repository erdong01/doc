package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var nums = []int{1}
	res := maxSubArray(nums)
	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

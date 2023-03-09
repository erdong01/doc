package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func searchRange(nums []int, target int) []int {

	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func lowerBound(nums []int, target int) int {

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
func searchRange2(nums []int, target int) []int {

	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lowerBound(nums, target+1) - 1
	return []int{start, end}
}

func TestXxx(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	res := searchRange2(nums, 8)
	fmt.Println(res)
}

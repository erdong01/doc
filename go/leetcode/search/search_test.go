package search_test

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := search(arr, 6)
	fmt.Println(i)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		fmt.Println(right, left, mid)
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 1
}

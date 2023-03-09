package leetcode

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	nums := []int{4, 9,5,8, 6, 7, 0, 1, 2}
	res := search(nums, 8)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	count := len(nums)
	if count == 0 {
		return -1
	}
	if count == 1 && nums[0] == target {
		return 0
	}
	var l, r = 0, count - 1
	for l <= r {
		mid := (l + r) / 2
		fmt.Println("mid", mid)
		if nums[mid] == target {
			return mid
		}
		fmt.Println("nums[0],nums[mid]", nums[0], nums[mid])
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[count-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

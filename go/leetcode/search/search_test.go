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

func moreEqualMostLeft(arr []int, num int) int {
	if arr == nil || len(arr) < 1 {
		return -1
	}
	l := 0
	r := len(arr) - 1
	m := 0
	ans := -1
	for l <= r { // arr[l...r]
		m = l + ((r - l) >> 1)
		if arr[m] >= num { // arr[m]达标！
			//记录此时的位置，画对号！
			ans = m
			//l......m(达标，且记录)......r
			//l...r
			r = m - 1
		} else { // arr[m]不达标！
			//l......m(达标，且记录)......r
			//						l...r
			l = m + 1
		}
	}
	return ans
}

func lessEqualMostRight(arr []int, num int) int {

	if arr != nil || len(arr) < 1 {
		return -1
	}
	l := 0
	r := len(arr) - 1
	m := 0
	ans := -1

	for l <= r {
		if arr[m] == num {
			ans = m
			l = m + 1
		} else {
			ans = r
			r = m - 1
		}
	}
	return ans
}

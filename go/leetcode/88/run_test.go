package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var m = 3
	var nums2 = []int{2, 5, 6}
	var n = 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, tail := m-1, n-1, m+n-1; i >= 0 || j >= 0; tail-- {
		var cur int
		if i == -1 {
			cur = nums2[j]
			j--
		} else if j == -1 {
			cur = nums1[i]
			i--
		} else if nums1[i] > nums2[j] {
			cur = nums1[i]
			i--
		} else {
			cur = nums2[j]
			j--
		}
		nums1[tail] = cur
	}
}
func merge1(nums1 []int, m int, nums2 []int, n int) {
	var result []int
	var i, j int
	for i < m || j < n {
		if j >= n || (i < m && nums1[i] <= nums2[j]) {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	var count = len(nums1)
	for k := 0; k < count; k++ {
		nums1[k] = result[k]
	}
}

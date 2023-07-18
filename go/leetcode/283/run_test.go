package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	a(1)
	a("1")
	var nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
func moveZeroes(nums []int) {
	var left, right, n = 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			fmt.Println("nums", nums)
			left++
		}
		right++
	}
}
func moveZeroes1(nums []int) {
	var n int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			n++
		}
	}

	for n < len(nums) {
		nums[n] = 0
		n++
	}
}

type T interface {
	int | string | bool
}

func a[t T](a t) {
	fmt.Println(a)
}

package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var a, b = 60, 13
	fmt.Println(a ^ b)

	var nums = []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	var k = 2
	res := numberOfSubarrays(nums, k)
	fmt.Println("res", res)
}

func numberOfSubarrays(nums []int, k int) int {
	var arr = make([]int, len(nums)+1)
	arr[0] = 1
	res, sum := 0, 0
	for _, num := range nums {
		sum += num & 1
		fmt.Println("sum", sum, num)
		arr[sum]++
		fmt.Println("arr", arr)
		if sum >= k {
			fmt.Println("sum ,k", sum, k, sum-k)
			res += arr[sum-k]
		}
	}
	return res

}

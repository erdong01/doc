package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}

func subsets(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return ans
}

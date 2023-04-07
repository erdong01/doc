package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var numbers = []int{2, 7, 11, 15}
	var target = 9
	res := twoSum(numbers, target)
	fmt.Println(res)
}

func twoSum(numbers []int, target int) []int {
	j := len(numbers) - 1
	for i := 0; i < len(numbers); i++ {
		for ; i < j && numbers[i]+numbers[j] > target; j-- {
		}
		if i < j && numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

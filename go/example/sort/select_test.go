package sort_test

import (
	"fmt"
	"testing"
)

//选择排序
func TestSelectSort(t *testing.T) {
	arr := [6]int{10, 34, 19, 100, 80, 99}
	SelectSort(&arr)
	fmt.Println(arr)
}

func SelectSort(arr *[6]int) {

	for j := 0; j < len(arr)-1; j++ {
		maxIndex := j
		max := arr[j]
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
	fmt.Println(arr)
}

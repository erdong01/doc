package preSum

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	arr := []int{2, 3, 1, 5, -2}
	sumArr := preSumArray(arr)
	fmt.Println(sumArr)

	//求 1到4的和
	sum := getSum(sumArr, 1, 4)
	fmt.Println("sum", sum)
}

func preSumArray(arr []int) (sumArr []int) {
	n := len(arr)
	sumArr = make([]int, n)
	sumArr[0] = arr[0]
	for i := 1; i < n; i++ {
		sumArr[i] = arr[i] + sumArr[i-1]
	}
	return
}

func getSum(arr []int, l int, r int) (sum int) {
	if r == 0 {
		return arr[0]
	}
	sum = arr[r] - arr[l-1]
	return
}

package tree

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	a := []int{2, 5}

	b := 5

	idx := searchInsert(a, b)

	fmt.Println(idx) //3
	idx2 := insertSearch(a, b)

	fmt.Println(idx2) //3
}

//二分查找算法
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1

}

//差值查找
func insertSearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	time := 0
	for low < high {
		time += 1
		mid := low + int((high-low)*(key-arr[low])/(arr[high]-arr[low]))
		if key < arr[mid] {
			high = mid - 1
		} else if key > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func TestFibonacciSearch(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5}
	var key int = 2
	index := fibonacciSearch(arr, key)
	fmt.Println(index)
}

//斐波那契查找
func fibonacciSearch(arr []int, key int) int {

	fibArr := make([]int, 0)
	for i := 0; i <= 36; i++ {
		fibArr = append(fibArr, fibonacci(i))
	}
	//确定待查找数组在斐波那契数列的位置
	k := 0
	n := len(arr)
	// 此处 n > fib[k]-1 也是别有深意的
	// 若n恰好是裴波那契数列上某一项，且要查找的元素正好在最后一位，此时必须将数组长度填充到数列下一项的数字
	for n > fibArr[k]-1 {
		k = k + 1
	}

	for i := n; i < fibArr[k]; i++ {
		arr = append(arr, 0)
	}

	low, high := 0, n-1
	for low <= high {
		mid := low + fibArr[k-1] - 1
		if key < arr[mid] {
			high = mid - 1
			k -= 1

		} else if key > arr[mid] {
			low = mid + 1
			k -= 2
		} else {
			if mid < n {
				return mid
			} else {
				return n - 1
			}
		}
	}
	return -1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	var fibarry = [3]int{0, 1, 0}
	for i := 2; i <= n; i++ {
		fibarry[2] = fibarry[0] + fibarry[1]
		fibarry[0] = fibarry[1]
		fibarry[1] = fibarry[2]
	}
	return fibarry[2]
}

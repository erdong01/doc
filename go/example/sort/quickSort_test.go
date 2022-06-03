package sort_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFindKthLargest(t *testing.T) {
	index := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	fmt.Println(index)
}

//215 快速排序
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func TestXxx2(t *testing.T) {
	arr := [6]int{3, 2, 1, 5, 6, 4}
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}

func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right

	pivot := array[(left+right)/2]

	for l < r {
		for array[l] < pivot {
			l++
		}
		for array[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		temp := array[l]
		array[l] = array[r]
		array[r] = temp
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}
}

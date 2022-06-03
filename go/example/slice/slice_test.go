package slice_test

import (
	"fmt"
	"testing"
	"time"
)

func TestForRange(t *testing.T) {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for i := range arr {
		newArr = append(newArr, &arr[i])

	}
	for _, v := range newArr {
		fmt.Println(*v)
	}

	var val []int
	val = append(val, 2000)
	appendTest(val)
	fmt.Println(val)
}

func appendTest(aa []int) {
	aa = append(aa, 2)
	for i := 0; i < 10; i++ {

		aa = append(aa, i)
	}

}

func TestAppend(t *testing.T) {
	s := make([]int, 5)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)

	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2)
}

func TestDDD(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func TestFunc(t *testing.T) {
	a := []int{7, 8, 9}
	fmt.Printf("%+v \n", a)
	ap(a)
	fmt.Printf("%+v \n", a)
	app(a)
	fmt.Printf("%+v \n", a)
}

func ap(a []int) {
	a = append(a, 10)
}

func app(a []int) {
	a[0] = 1
}

func TestForFor(t *testing.T) {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}

func TestAssignment(t *testing.T) {
	var a1 uint = 1
	var b1 uint = 2
	fmt.Println(a1 - b1)

	var a = []int{1, 2, 3, 4, 5}
	var r [5]int
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r =", r)
	fmt.Println("a =", a)
	// r 1 12 13 45
	// a 1 12 13 4 5
}

func TestJh(t *testing.T) {
	x := []int{1, 2, 3}
	y := x[:2]

	y = append(y, 50)
	y = append(y, 60)
	fmt.Printf("x = %v \n", x) // 1 2  50 60
	y[0] = 20
	fmt.Printf("x = %v \n", x) // 1 2 50 60
}

type Obj struct {
	id int
}

func createObjes(n int) []Obj {

	if n%2 == 0 {

	} else {
		n++
	}
	refs := make([]Obj, n)
	for i := 0; i < n; i++ {
		obj := Obj{id: i}
		refs[i] = obj
		// obj2 := Obj{id: i + 1}
		// refs[i+1] = obj2
	}

	return refs
}

func TestCreateObjes(t *testing.T) {
	start := time.Now()
	createObjes(50000000)

	end := time.Now()
	fmt.Println(end.Sub(start))
}

func TestJq(t *testing.T) {
	var arr = []int{1, 2, 3, 4}
	arr = append(arr[:0], arr[2:]...)
	fmt.Println(arr)
}

func TestForAdd(t *testing.T) {
	var arr = []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, i)
	}
	start := time.Now()
	for i := 0; i < 100000; i++ {
		// arr = removeCopy(arr, 0)
		arr = append(arr[:0], arr[0+1:]...)
	}
	end := time.Now()
	fmt.Println(end.Sub(start), arr)
}
func remove(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func removeCopy(data []int, index int) []int {

	copy(data[index:], data[index+1:])
	return data[:len(data)-1]
}

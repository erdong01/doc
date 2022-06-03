package map_test

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := make(map[int]*int)
	for key := range slice {
		m[key] = &slice[key]
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

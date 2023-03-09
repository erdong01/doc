package string_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestCopy(t *testing.T) {
	s0 := "煎鱼是脑子进水了"
	s1 := Clone(s0)
	// s1 = s1[:3]
	fmt.Println(s0)
	fmt.Println(s1)
	s0h := (*reflect.StringHeader)(unsafe.Pointer(&s0))
	s1h := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Println(s0h.Len, s1h.Len)
	fmt.Println(s0h.Data, s1h.Data)
}

func Clone(s string) string {
	if len(s) == 0 {
		return ""
	}
	b := make([]byte, len(s))
	copy(b, s)
	return *(*string)(unsafe.Pointer(&b))
}

func reverse(str string) string {
	s := []rune(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func TestXxx(t *testing.T) {
	str := "hello"
	a := []rune(str)
	a[0] = 'x'
	fmt.Println(str[0])
	fmt.Println(string(a))
}

package t_test

import "testing"

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~complex64 | ~complex128
}

func Neg[T signed](n T) T {
	return -n
}

func TestNeg(t *testing.T) {}

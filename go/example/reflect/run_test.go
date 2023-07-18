package reflect_test

import "testing"

type I interface {
	foo()
	Bar()
}
type T struct{}

func (T) foo() {}
func (T) Bar() {}

func TestGo1011(tt *testing.T) {

}

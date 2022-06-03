package factorymethod

import "testing"

func compute(factory OperatorFactory, a, b int) int {

	of := factory.Create()
	of.SetA(a)
	of.SetB(b)
	return of.Result()
}

func TestXxx(t *testing.T) {
	var factory OperatorFactory //工厂接口

	factory =PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}
}

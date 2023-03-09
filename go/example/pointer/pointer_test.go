package pointer_test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	a := 20
	b := &a
	fmt.Printf("type of b:%T \n", b)

	c := *b
	fmt.Printf("type of c:%T \n", c)
	fmt.Printf("value of c:%v \n", c)
}

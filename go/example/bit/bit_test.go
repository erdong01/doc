package bit

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var a uint = 60
	var b uint = 13
	var c uint

	c = a & b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	fmt.Printf("第一行 - c 的值为 %d\n", b<<2)
}

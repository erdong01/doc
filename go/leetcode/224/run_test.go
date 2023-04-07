package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	ans := calculate("2-2 + 2 ")
	fmt.Println(ans)
}

func calculate(s string) (ans int) {
	ops := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			fmt.Println("sign", sign)
			i++
		case '-':
			sign = -ops[len(ops)-1]
			fmt.Println("sign", sign)
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0

			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
				fmt.Println("num", num)
			}
			ans += sign * num
		}
	}
	return
}

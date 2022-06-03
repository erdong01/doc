package select_test

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {

	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			fmt.Println("case1")
		case <-ch:
			fmt.Println("case2")
		}
	}
}
func TestChan2(t *testing.T) {

	ch := make(chan int, 2)
	go func() {
		ch <- 1
	}()
	select {
	case <-ch:
		fmt.Println("case")

	}
}

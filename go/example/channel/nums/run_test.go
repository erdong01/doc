package nums_test

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for n := range ch1 {
			fmt.Printf("channel-1 print %d \n", n)
			ch2 <- n + 1

			if n == 9 {
				break
			}
		}
		close(ch1)
	}()

	go func() {
		for n := range ch2 {
			fmt.Printf("channel-2 print %d \n", n)
			if n == 10 {
				break
			}
			ch1 <- n + 1
		}
		close(ch2)
	}()
	ch1 <- 1
	time.Sleep(time.Second)
}

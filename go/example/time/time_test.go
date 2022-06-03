package time_test

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	tickerIn := time.NewTicker(time.Second * 3)
	tickerOut := time.NewTicker(time.Second * 6)
	for {
		select {
		case <-tickerIn.C:
			fmt.Println("3秒IN")
		case <-tickerOut.C:
			fmt.Println("5秒OUT")
		}
	}
}

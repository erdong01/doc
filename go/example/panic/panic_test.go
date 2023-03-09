package panic

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()
	select {}
}

func proc() {
	panic("ok")
}


package goroutine

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"testing"
	"time"
)

func cat(fishCH, catCH chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			fmt.Println("cat")
			catCH <- struct{}{}
			<-fishCH
		}
		wg.Done()
	}()
}
func dog(catCH, dogCH chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			<-catCH
			fmt.Println("dog")
			dogCH <- struct{}{}
		}
		wg.Done()
	}()
}

func fish(dogCH, fishCH chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			<-dogCH
			fmt.Println("fish")
			time.Sleep(time.Second)
			fishCH <- struct{}{}
		}
		wg.Done()
	}()
}
func TestXxx(t *testing.T) {
	catCH := make(chan struct{})
	dogCH := make(chan struct{})
	fishCH := make(chan struct{})
	wg := sync.WaitGroup{}
	cat(fishCH, catCH, &wg)
	dog(catCH, dogCH, &wg)
	fish(dogCH, fishCH, &wg)
	wg.Wait()
}

func TestTrace(t *testing.T) {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	trace.Start(f)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello GMP")
	trace.Stop()
}

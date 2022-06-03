package channel_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestA(t *testing.T) {
	var wg sync.WaitGroup
	var catCounter uint32
	var dogCounter uint32
	var fishCounter uint32
	catCh := make(chan struct{}, 1)
	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)
	wg.Add(3)

	go cat(&wg, catCounter, catCh, dogCh)
	go dog(&wg, dogCounter, dogCh, fishCh)
	go fish(&wg, fishCounter, fishCh, catCh)
	catCh <- struct{}{}
	wg.Wait()
}

func cat(wg *sync.WaitGroup, counter uint32, catCh chan struct{}, dogCh chan struct{}) {
	for {
		if counter >= uint32(100) {
			wg.Done()
			return
		}
		<-catCh
		fmt.Println("cat")
		atomic.AddUint32(&counter, 1)
		dogCh <- struct{}{}
	}
}
func dog(wg *sync.WaitGroup, counter uint32, dogCh chan struct{}, fishCh chan struct{}) {
	for {
		if counter >= uint32(100) {
			wg.Done()
			return
		}
		<-dogCh
		fmt.Println("dog")
		atomic.AddUint32(&counter, 1)
		fishCh <- struct{}{}
	}
}
func fish(wg *sync.WaitGroup, counter uint32, fishCh chan struct{}, catCh chan struct{}) {

	for {
		if counter >= uint32(100) {
			wg.Done()
			return
		}
		<-fishCh
		fmt.Println("fish")
		atomic.AddUint32(&counter, 1)
		catCh <- struct{}{}
	}
}

func TestCat(t *testing.T) {
	var count = 100
	var wg sync.WaitGroup
	wg.Add(count)
	ch := make(chan struct{})
	c := make(chan struct{}, 1)
	for i := 0; i < count; i++ {
		c <- struct{}{}
		next := cat1(ch)
		next = dog1(next)
		fish1(next, c)
		ch <- struct{}{}
		wg.Done()
	}
	wg.Wait()
}

func cat1(ch <-chan struct{}) <-chan struct{} {
	out := make(chan struct{})
	go func() {
		defer close(out)
		<-ch
		fmt.Println("cat")

	}()
	return out
}

func dog1(ch <-chan struct{}) <-chan struct{} {
	out := make(chan struct{})
	go func() {
		defer close(out)
		<-ch
		fmt.Println("dog")

	}()
	return out
}

func fish1(ch <-chan struct{}, c <-chan struct{}) {
	go func() {
		<-ch
		fmt.Println("fish")
		<-c

	}()
	return
}

func TestJtdy(t *testing.T) {

}

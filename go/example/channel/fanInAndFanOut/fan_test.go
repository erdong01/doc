package faninandfanout

import (
	"fmt"
	"sync"
	"testing"
)

func TestMain(t *testing.T) {
	coms := buy(10)

	phones1 := build(coms)
	phones2 := build(coms)
	phones3 := build(coms)
	phones := merge(phones1,phones2,phones3)
	packs := pack(phones) 
	for p := range packs {
		fmt.Println(p)
	}
}

func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("零件", i)
		}
	}()
	return out
}

func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包（" + c + ")"
		}
	}()
	return out
}

func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}
	wg.Add(len(ins))
	for _, cs := range ins {
		go p(cs)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

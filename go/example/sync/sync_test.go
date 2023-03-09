package sync_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var pool = sync.Pool{
	New: func() interface{} {
		return "123"
	},
}

func TestPool(tt *testing.T) {

	t := pool.Get().(string)
	fmt.Println(t)
	t3 := pool.Get().(string)
	fmt.Println(t3)
	pool.Put("321")
	pool.Put("321")
	pool.Put("321")
	pool.Put("321")

	runtime.GC()
	time.Sleep(1 * time.Second)

	t2 := pool.Get().(string)
	fmt.Println(t2)

	// runtime.GC()
	// time.Sleep(1 * time.Second)

	t2 = pool.Get().(string)
	fmt.Println(t2)
	t2 = pool.Get().(string)
	fmt.Println(t2)
	t2 = pool.Get().(string)
	fmt.Println(t2)
	t2 = pool.Get().(string)
	fmt.Println(t2)
	t2 = pool.Get().(string)
	fmt.Println(t2)

	t2 = pool.Get().(string)
	fmt.Println(t2)

	t2 = pool.Get().(string)
	fmt.Println(t2)

	t2 = pool.Get().(string)
	fmt.Println(t2)

	t2 = pool.Get().(string)
	fmt.Println(t2)
}

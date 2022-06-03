package channel_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

type T struct {
	a1 int64
	a  int32
}

func TestXxx(t *testing.T) {
	s := []string{"1", "2", "3"}
	s2 := []string{"1"}
	fmt.Println(unsafe.Sizeof(s))  // 24
	fmt.Println(unsafe.Sizeof(s2)) // 24
	var t1 T
	fmt.Println(unsafe.Sizeof(t1))
	// s := []int{7, 2, 8, -9, 4, 0}
	// c := make(chan int)
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c
	// fmt.Println(x, y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func TestChannleVar(t *testing.T) {
	var chan1 chan struct{}
	chan1 = make(chan struct{})
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-chan1)

	}()

	chan1 <- struct{}{}

	fmt.Println("aa")

	ch <- 1
	fmt.Println("bb")
	time.Sleep(time.Second)
}

func TestTimeout(t *testing.T) {
	result := make(chan string)
	timeout := time.After(3 * time.Second)
	go func() {
		time.Sleep(1 * time.Second)
		close(result)
	}()

	for {
		select {
		case v, ok := <-result:
			if !ok {
				fmt.Println("通道已关闭")
				return
			}
			fmt.Println("v", v)
		case <-timeout:
			fmt.Println("网络访问超时了")
			return
		default:
			fmt.Println("等待。。。")
			time.Sleep(1 * time.Second)
		}
	}
}

var wg sync.WaitGroup

func TestWorker(t *testing.T) {
	requestCount := 10
	worker := 3
	ch := make(chan struct{}, worker)

	for i := 0; i < requestCount; i++ {
		wg.Add(i)
		ch <- struct{}{}
		go Read(ch, i)
	}
	wg.Wait()
}

func Read(ch chan struct{}, i int) {
	fmt.Println(i)
	<-ch
	wg.Done()
}

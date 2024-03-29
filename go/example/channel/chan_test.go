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

var M map[int]int
var CM chan int

func RunCM() {
	for {
		select {
		case m := <-CM:
			fmt.Println("m", m)
			M[m] = m
		}
	}
}

func TestXxxx(t *testing.T) {
	M = make(map[int]int)
	CM = make(chan int)
	go RunCM()
	for i := 0; i < 1000; i++ {
		CM <- i
	}
	fmt.Println(len(M))
}

func TestP(t *testing.T) {
	// var a int
	// var p *int
	// p = &a

	// a = 10

	// fmt.Println("*p = ", *p)
	// fmt.Println("a = ", a)

	// *p = 100 //解引用
	// fmt.Println("*p = ", *p)
	// fmt.Println("a = ", a)

	var ss *string
	ss = new(string)
	*ss = "1"
	fmt.Printf("%q \n", *ss)
}

func TestAAAA(t *testing.T) {
	ch := make(chan string, 4)
	fmt.Println(ch, unsafe.Sizeof(ch))
	go sendTask(ch)
	go receiveTask(ch)
	time.Sleep(1 * time.Second)
}

func sendTask(ch chan string) {
	taskList := []string{"this", "is", "a", "demo"}
	for _, task := range taskList {
		ch <- task
	}
}

func receiveTask(ch chan string) {
	for {
		task := <-ch
		fmt.Println("received:", task)
	}
}

func TestChan(t *testing.T) {

	chan1 := make(chan struct{}, 1)
	chan2 := make(chan struct{}, 1)
	chan3 := make(chan struct{}, 1)
	wg.Add(3)
	chan1 <- struct{}{}
	start := time.Now().Unix()
	go print("1", chan1, chan2)
	go print("2", chan2, chan3)
	go print("3", chan3, chan1)
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start)
}

func print(goroutine string, inputchan chan struct{}, outchan chan struct{}) {
	time.Sleep(1 * time.Second)
	select {
	case <-inputchan:
		fmt.Printf("%s", goroutine)
		outchan <- struct{}{}
	}
	wg.Done()
}

package map_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFor(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		m[key] = &val
		fmt.Println(&val)
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func TestMaxSubArrayLen(t *testing.T) {
	var arr = []int{1, 2, 5, 6, -9, 7, 1, 1}
	l := maxSubArrayLen(arr, 5)
	fmt.Println(l)
}

func maxSubArrayLen(nums []int, k int) int {
	ans := 0

	sum := 0

	sumMap := map[int]int{
		0: -1,
	}

	for i, num := range nums {
		sum += num
		pre, ok := sumMap[sum-k]
		if ok && ans < i-pre {
			ans = i - pre
		}
		if _, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		}
	}

	return ans
}

var Data map[int]string
var DataMuetx sync.Mutex
var DataChan chan Dataer

type Dataer struct {
	State int8
	Key   int
	Value string
}

func TestXxx(t *testing.T) {
	DataChan = make(chan Dataer)
	Data = make(map[int]string)
	go ChanData(DataChan)
	var wg sync.WaitGroup
	var i = 200
	wg.Add(i)
	start := time.Now()
	for i > 0 {
		i--
		go func() {
			var dataer Dataer
			dataer.State = 1
			dataer.Key = 111
			dataer.Value = "52384572309573029573298573902573092570325730257302570235702502345723057032945703245723045"
			DataChan <- dataer
			var dataer2 Dataer
			dataer2.State = 2
			dataer2.Key = 111
			DataChan <- dataer2
			dataer2 = <-DataChan
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func DataWriteMeutex(key int, value string) {
	DataMuetx.Lock()
	defer DataMuetx.Unlock()
	Data[key] = value
}
func DataReadMeutex(key int) (value string) {
	DataMuetx.Lock()
	defer DataMuetx.Unlock()
	value = Data[key]
	return value
}

func ChanData(DataChan chan Dataer) {
	for {
		data := <-DataChan
		if data.State == 1 {
			Data[data.Key] = data.Value
		} else {
			data.Value = Data[data.Key]
			DataChan <- data
		}
	}
}

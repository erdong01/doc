package memory

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var memory *MemoryTest
var one sync.Once

type MemoryTest struct {
	Lock sync.RWMutex
	Arr  []int
}

func New() *MemoryTest {
	one.Do(func() {
		memory = &MemoryTest{}
	})
	return memory
}

func (m *MemoryTest) AGo(i int) {
	m.Lock.Lock()
	m.Arr = append(m.Arr, i)
	m.Lock.Unlock()
}

func (m *MemoryTest) BGo(i int) {
	m.Lock.Lock()
	m.Arr = append(m.Arr, i)
	m.Lock.Unlock()
}

func (m *MemoryTest) CGo(i int) {
	m.Lock.Unlock()
	m.Arr = append(m.Arr, i)
	m.Lock.Unlock()
}

func (m *MemoryTest) DGo(i int) {
	m.Lock.Unlock()
	m.Arr = append(m.Arr, i)
	m.Lock.Unlock()
}

func TestXxx(t *testing.T) {
	m := New()

	for i := 0; i < 100000; i++ {
		go func(v int) {
			m.AGo(v)
		}(i)
		go func(v int) {
			m.BGo(v + 1)
		}(i)
		// go func(v int) {
		// 	m.CGo(v + 2)
		// }(i)
		// go func(v int) {
		// 	m.DGo(v)
		// }(i)

	}

	time.Sleep(time.Second * 1)
	fmt.Println(len(m.Arr))
}

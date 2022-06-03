package singleton

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(parCount)                    //预计使用100个协程并发,计数器为100
	instances := [parCount]*Singleton{} //初始化100个单例数组
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		log.Printf("i=%d", i)
		if instances[i] != instances[i-1] { //对比相邻的两个单例是否相等
			t.Fatal("instance is not equal")
		}
	}
	log.Printf("执行完毕")
	fmt.Println(instances)
}

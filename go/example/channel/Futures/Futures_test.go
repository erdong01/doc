package futures

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	vegetablesCh := washVegetables() //洗菜
	waterCh := boliWater()           //烧水
	fmt.Println("已经安排好洗菜和烧水了，我先开一局")
	time.Sleep(2 * time.Second)
	fmt.Println("要做火锅了，看看菜和水好了吗")
	vegetables := <-vegetablesCh
	water := <-waterCh
	fmt.Println("准备好了，可以做火锅了:", vegetables, water)

}

func washVegetables() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "洗好的菜"
	}()
	return vegetables
}

//烧水
func boliWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "烧开的水"
	}()
	return water
}

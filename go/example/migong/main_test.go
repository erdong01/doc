package migong

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	//元素是1 表示 墙
	//元素是0 表示 未知
	//元素是2 表示 可以通行
	//元素是3 表示 走过 但是路不通
	var myMap [8][7]int

	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	SetWay(&myMap, 1, 1)

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}

func SetWay(myMap *[8][7]int, i, j int) bool {

	//什么情况下找到出路

	//map[6][5] ==2 表示找到出路

	if myMap[6][5] == 2 {
		fmt.Println("路找到了")
		return true
	}
	if myMap[i][j] == 0 {
		myMap[i][j] = 2
		if SetWay(myMap, i+1, j) {
			return true
		} else if SetWay(myMap, i, j+1) {
			return true
		} else if SetWay(myMap, i-1, j) {
			return true
		} else if SetWay(myMap, i, j-1) {
			return true
		} else {
			//死路
			myMap[i][j] = 3
			return false
		}

	} else {
		return false
	}

	return false
}

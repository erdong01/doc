package main

import (
	"fmt"
	"log"
	"sort"
)

var ruiJinmahjongArr = []int{
	101, 102, 103, 104, 105, 106, 107, 108, 109, //#万
	101, 102, 103, 104, 105, 106, 107, 108, 109,
	101, 102, 103, 104, 105, 106, 107, 108, 109,
	101, 102, 103, 104, 105, 106, 107, 108, 109,
	201, 202, 203, 204, 205, 206, 207, 208, 209, //#饼
	201, 202, 203, 204, 205, 206, 207, 208, 209,
	201, 202, 203, 204, 205, 206, 207, 208, 209,
	201, 202, 203, 204, 205, 206, 207, 208, 209,
	301, 302, 303, 304, 305, 306, 307, 308, 309, //#条
	301, 302, 303, 304, 305, 306, 307, 308, 309,
	301, 302, 303, 304, 305, 306, 307, 308, 309,
	301, 302, 303, 304, 305, 306, 307, 308, 309,
	401, 402, 403, 404, 405, 406, 407, //# 东 西 南 北 中 发 白
	401, 402, 403, 404, 405, 406, 407,
	401, 402, 403, 404, 405, 406, 407,
	401, 402, 403, 404, 405, 406, 407,
	501, 502, 503, 504, //花(春夏秋冬)
}

const (
	GroupTypeSingle byte = iota //单张
	GroupTypeDouble             //对子
	GroupTypePenZi              //碰子
	GroupTypeShunZi             //顺子
)

type mahjong struct {
	DuiDui [][]int
	Sunzi  [][]int //顺子
	Carve  [][]int //刻子
	one    []int
}

func isHu(arr []int) bool {
	mjArr := append([]int{}, arr...)
	//检测牌数量
	if len(mjArr)%3 != 2 {
		log.Println("牌数量不符合3n+2")
		return false
	}
	var mahjong mahjong
	Sort(arr)

	var target int
	var newArr []int
	target = arr[0]
	var isSunzi bool

	var sunziCount int = 0

	for key := range arr {
		if target+sunziCount != arr[key] {
			if len(newArr) == 1 {
				mahjong.one = append(mahjong.one, newArr...)
			} else if len(newArr) == 2 && isSunzi == false {
				mahjong.DuiDui = append(mahjong.DuiDui, newArr)
			} else if len(newArr) == 3 && isSunzi == true {
				isSunzi = false
				mahjong.Sunzi = append(mahjong.Sunzi, newArr)
			} else if len(newArr) < 3 && isSunzi == true {
				mahjong.one = append(mahjong.one, newArr...)
			} else if len(newArr) == 3 || len(newArr) == 4 {
				mahjong.Carve = append(mahjong.Carve, newArr)
			}
			newArr = []int{}
			target = arr[key]
			sunziCount = 0
		}
		if target+sunziCount == arr[key] {
			newArr = append(newArr, arr[key])
		}
		if isSunzi == true {
			sunziCount++
		}
		if len(arr) > key+1 && target+1 == arr[key+1] && len(newArr) == 1 {
			fmt.Println(target+sunziCount, arr[key], newArr, "-------")
			isSunzi = true
			sunziCount++
		}

		if len(arr) == key+1 {
			// fmt.Println(newArr)
			if len(newArr) == 1 {
				mahjong.one = append(mahjong.one, newArr...)
			} else if len(newArr) == 2 && isSunzi == false {
				mahjong.DuiDui = append(mahjong.DuiDui, newArr)
			} else if len(newArr) == 3 && isSunzi == true {
				isSunzi = false
				mahjong.Sunzi = append(mahjong.Sunzi, newArr)
			} else if len(newArr) < 3 && isSunzi == true {
				mahjong.one = append(mahjong.one, newArr...)
			} else if len(newArr) == 3 || len(newArr) == 4 {
				mahjong.Carve = append(mahjong.Carve, newArr)
			}
		}
	}
	duiduiCount := len(mahjong.DuiDui)
	oneCount := len(mahjong.one)
	carveCount := len(mahjong.Carve)
	sunzicount := len(mahjong.Sunzi)
	fmt.Println(duiduiCount, carveCount, oneCount, sunzicount)
	//七对
	if duiduiCount == 7 {
		return true
	} else if duiduiCount >= 4 {

		ThreeDuiToSunzi(&mahjong)
		duiduiCount = len(mahjong.DuiDui)
	}

	if oneCount == 0 && duiduiCount == 1 {
		return true
	}
	return false

}

func ThreeDuiToSunzi(mahjong *mahjong) {
	var target int
	target = mahjong.DuiDui[0][0]
	var sunzi1 []int
	var sunzi2 []int
	for duiKey := 0; duiKey < len(mahjong.DuiDui); duiKey++ {
		dui := mahjong.DuiDui[duiKey][0]
		if target == dui {
			target++
			sunzi1 = append(sunzi1, dui)
			sunzi2 = append(sunzi2, dui)
		} else if len(mahjong.DuiDui) > duiKey+1 {
			target = mahjong.DuiDui[duiKey+1][0]
			sunzi1 = []int{}
			sunzi2 = []int{}
		}
		if len(sunzi1) == 3 {
			mahjong.Sunzi = append(mahjong.Sunzi, sunzi1)
			mahjong.Sunzi = append(mahjong.Sunzi, sunzi2)
			mahjong.DuiDui = append(mahjong.DuiDui[:duiKey-2], mahjong.DuiDui[duiKey+1:]...)
			sunzi1 = []int{}
			sunzi2 = []int{}
			duiKey = duiKey - 3
		}

	}
}

func Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	sort.IntsAreSorted(arr)
	sort.Ints(arr)
	return arr
}

func main() {
	var arr = []int{101, 101, 102, 102, 103, 103, 201, 202, 203, 303, 303, 303, 402, 402}
	b := isHu(arr)
	fmt.Println(b)
}

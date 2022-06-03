package main

import "fmt"

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("arr %d ", v2)
		}
		fmt.Println()
	}
	var sparseArr []ValNode
	sparseArr = append(sparseArr, ValNode{
		row: 0,
		col: 0,
		val: 0,
	})
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				sparseArr = append(sparseArr, ValNode{
					row: i,
					col: j,
					val: v2,
				})
			}
		}
	}

	for i, valNode := range sparseArr {
		fmt.Println(i, valNode)
	}
}

type ValNode struct {
	row int
	col int
	val int
}

package main

import (
	"fmt"
)

func main() {
	var n = []int{1,2,3,4,5,6,7,8}
	fmt.Println(n)
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])
	n[2] = 100
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
	}
	fmt.Println(board[0][1])
	n = append(n, 100)
	fmt.Println(n)

}

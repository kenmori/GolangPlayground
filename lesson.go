package main

import (
	"fmt"
)

func main() {
	// 一番少ない値を出力するコード
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var min [1]int
	for i, v := range l {
		fmt.Println(i, v)
		if min[0] > v {
			min[0] = v
		}
	}
	fmt.Println("min is", min[0])
	// 合計を出力するコード
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	var amont int
	for _, v := range m {
		amont = +v
	}
	fmt.Println(amont)
}

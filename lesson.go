package main

import (
	"fmt"
)

func main() {
	// 一番少ない値を出力するコード
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var min int
	for i, num := range l {
		if i == 0 {
			min = num
			continue
		}
		if min >= num {
			min = num
		}
	}
	fmt.Println("min is", min)
	// 合計を出力するコード
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	sum := 0
	for _, v := range m {
		sum += v // 足す時
	}
	fmt.Println(sum)
}

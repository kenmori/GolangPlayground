package main

import (
	"fmt"
)

func main() {
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%d", len(n), cap(n), n)
	// m := map[string]int{"apple": 1, "orenge": 2}
	c := make([]int, 5) // この2つのcの作られ方違い
	// c := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	// m := map[string]int{"apple": 100}
	// v, ok := m["apple"]
	// m2 := make(map[string]int) // メモリ上に作る
	// m2["PC"] = 5000
	// fmt.Println(m2) // map[PC:5000]
	var m3 map[string]int // 宣言しただけ。初期化はしていない。 nilが入る
	// m3["PC"] = 5000;  // Error がメモリ上に入れるmapがないのでErrorになる
	if m3 == nil {
		fmt.Println("nil m2")
	}
	// fmt.Println(m3)
	var s []int // varで宣言した時はスライスの時でもmapのときでも宣言しただけの状態はnil
	if s == nil {
		fmt.Println("Nil")
	}
}

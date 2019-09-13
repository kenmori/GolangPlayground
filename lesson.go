package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx) // %f float表記
	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	// i := int(s); // 14にはならない
	// コンバートするには
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v", i, i) // Atoi アスキートゥインテジャー
}

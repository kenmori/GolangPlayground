package main

import (
	"fmt"
)

func main() {
	var b = []int{100}
	b = append(b, 300)
	fmt.Println(b)
}

package main

import (
	"fmt"
)

type Vertex struct{
	X,  Y int
}

func (v Vertex) Area () int { // 値レシーバー
	return v.X * v.Y
}
func (v *Vertex) Scale (i int) { // ポイントレシーバー。Vertexの値を書き換えられる
	v.X = v.X * i
	v.Y = v.Y * i
}
func New (x, y int) *Vertex {
	return &Vertex{x, y}
}
func main() {
	// v := Vertex{3, 4}
	// fmt.Println(v.Area()) // .でvのメソッドになる
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}


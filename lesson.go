package main

import (
	"fmt"
)

type Vertex struct{
	x, y int
}
type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D () int { // 値レシーバー
	return v.x * v.y * v.z
}
func (v *Vertex3D) Scale3D (i int) { // ポイントレシーバー。Vertexの値を書き換えられる
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}
func New (x, y, z int) *Vertex3D {
	fmt.Println(Vertex3D{Vertex{x, y}, z})
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	// v := Vertex{3, 4}
	// fmt.Println(v.Area()) // .でvのメソッドになる
	v := New(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D())
}


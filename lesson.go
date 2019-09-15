package main

import (
	"fmt"
)

type Vertex struct{
	X int
	Y int
	S string
}

func changeVertext(v Vertex){
	v.X = 1000
}
func changeVertext2(v *Vertex){ // *がstructなら参照している先がかわることを理解する
	// (*v).X = 1000 structの場合はわざわざこのように書かないでいい // 実態に対して変更を行なっている
	v.X = 1000
}

func main() {
	v1 := Vertex{1, 2, "test"}
	changeVertext(v1) // 値渡しなのでコピーされた変数に1000が入っている
	fmt.Println(v1)
	v0 := &Vertex{1, 2, "test"}
	changeVertext2(v0) // 参照渡しなので書き換えられる
	// fmt.Println(v0)
	fmt.Println(*v0) // 実態を見る

	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "text"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println("%T %v\n", v4, v4)

	var v5 Vertex  // v4と一緒の意味
	fmt.Printf("%T %v\n", v5, v5) // 初期値が入る

	v6 := new(Vertex) // pointerを返す
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)
	
	s := make([]int, 0) // sliceやmapを書く時はmakeを使うことが多い
	fmt.Printf("%T %v\n", s, s)
	// s := []int{}

}


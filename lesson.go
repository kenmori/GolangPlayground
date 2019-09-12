package main

import (
    "fmt"
)
var i int = 2
func foo() {
    xi := 1
    xi = 2 //そのまま代入 xi := 2宣言し直すとエラー。 var i = 2としてもエラー。関数外で宣言されているから
    xf64 := 1.2
    xs := "test"
    xt, xf := true, false
    fmt.Println(xi, xf64, xs, xt, xf) // lnは自動的に改行いれてくれる
    fmt.Printf("%T\n", xf64) // 型を調べる
    fmt.Printf("%T\n", xf64) // 型を調べる
}
func main() {
    // var i int = 1
    // var f64 float64 = 1.2
    // var s string = "test"
    // var t, f bool = true, false

    // パレンティスに一緒に宣言することでvar重複をなくす
    // var (
    //     i int = 1
    //     f64 float64 = 1.2
    //     s string = "test"
    //     t, f bool = true, false
    // )
    // 初期化しないとそれぞれdefualt値が入る
    // var (
    //     i int
    //     f64 float64
    //     s string
    //     t, f bool
    // )
    // 型を明示的に宣言したい場合 varをつかう
    // var xf32 float32 = 1.2
    // 関数内でしか宣言できない変数ショートカット(varは関数外でも宣言できる)
    foo()
}
// go doc fmt Println
// go doc os/user Current
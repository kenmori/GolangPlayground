package main

import (
	"fmt"
)

func do (i interface{}){ // どんな型でも受け取るので // interfaceのものの型を変更するのがタイプアサーション
	// ss := i.(string)
	// ii := i.(int) // ここでタイプアサーションしなくてはいけない
	// ii *= 2
	// fmt.Println(ii)
	// fmt.Println(ss + "!")
	switch v := i.(type) {
		case int:
			fmt.Printf("%T %d\n", v, v)
		case string:
			fmt.Printf("%T %s\n", v, v)
		default:
			fmt.Printf("i don know")

	}
	var a int = 10
	aa := float64(10) // タイプコンバージョン。interfaceでないのでタイプアサーションではない
	fmt.Println(a, aa)

}
func main() {
	// var i interface{} = 10 //do(10)と同じ意味。10を代入したからといってiはintegerではないので。そのご受け取り側でType Assertionする必要がある
	// do(i)
	do(10)
	do("Mike")
	do(true)	
}


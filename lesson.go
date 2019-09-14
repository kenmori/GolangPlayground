package main

import (
	"fmt"
)

func one(x *int) { // アドレスの中身を型(ポイント型)にしている アスタリスク型
	*x = 1 // xにはあたいは入れられないので デリファレンス(*x)にして値を入れる。中身に1をいれている
}
func main() {
	var n int = 100
	one(&n)           //　アドレス型を渡している アンンパサンド ポイント型にはアドレス型を渡さないといけない
	fmt.Println(n)    // 1 中身が書き換わった
	fmt.Println(&*&n) // メモリのアドレスが出力される

	fmt.Println(n)
	fmt.Println(&n)
	var p *int = &n // integerのポイント型 16真数 4バイトなので4ずつ増えている *intはアドレスが入る実態の型(ポイント型)
	fmt.Println(p)
	fmt.Println(*p) // ポイント型のpに入っているアドレスが指しているメモリの値を指している
	// ＆はアドレス
	// * はアドレスの実態を参照している
}

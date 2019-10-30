package main

import (
	"fmt"
	"sync"
)

// ゴルーチン 並列処理をする
func goroutine (str string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}
	wg.Done() // Add(1)なら1回Doneしないといけない
}
func normal (s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1) // 1つのdoneされるまで待つ
	go goroutine("world", &wg)
	normal("hello")
	wg.Wait() //待つ
}


// goroutineが終わる前に処理が終わってしまうのを防ぐためにsync.WaitGroupがある
// wg.Add(1) // 一つの並列処理があるとことを伝える
// goroutineに&wgを渡す


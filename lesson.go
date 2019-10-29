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


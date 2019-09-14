package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatalln("Error")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data) // 左に一つでもinitializeできるものがあれば定義済みのerrは初期化できる, := short decralation。errは以前のerrを上書きされている
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))
	if err = os.Chdir("test"); err != nil { // 返値が1つしかなくてifの中でしか使わない場合1lineで書ける。ここでのerrをinitailzeしようとしるとエラーになる。ここではoverrideしていて、意味合いが違う
		log.Fatalln("Error")
	}
}

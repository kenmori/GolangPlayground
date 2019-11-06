package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Hello, playground")
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content))
	if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil { // errしか返さないので1lineで書ける
		log.Fatalln(err)
	}
	r := bytes.NewBuffer([]byte("abc")) //　バイトにする
	content2, _ := ioutil.ReadAll(r)    // errorハンドルなしなので_にしている。バイトを読み込む
	fmt.Println(string(content2))       // stringをバイトにデータコンバージョン
}

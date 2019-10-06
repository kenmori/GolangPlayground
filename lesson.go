package main

import (
	"fmt"
)

func thirdPartyConnectDB() {
	panic("Unable to connect database!") // panicはなるべく使わないことがGolangでは推奨されている errorhandlingするべき
}
func save() {
	defer func() { // saveが終わったら実行sれる。panicで強制終了されるのでそのrecoverをしてる
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB() // deferより上に書くと強制終了してrecoverまで到達しない
}

func main() {
	save()
	fmt.Println("OK?")
}

package main

import (
    "fmt"
    "strings"
)
func main() {
    fmt.Println("Hello World")
    fmt.Println("Hello" + "World")
    fmt.Println(string("HelloWorld"[0]))
    var s string = "Hello World"
    s = strings.Replace(s, "H", "W", 1) // 最初の1つめ
    fmt.Println(s)
    fmt.Println(`
        Test
                Test
    Test
    `)
    fmt.Println("\"")
    fmt.Println(`"`)
}

package main

import (
    "fmt"
    "os/user"
    "time"
)

func main() {
    fmt.Println("fafa", time.Now())
    fmt.Println(user.Current())
}
// go doc fmt Println
// go doc os/user Current
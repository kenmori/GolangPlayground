package main

import (
    "fmt"
)
func main() {
    t, f := true, false
    fmt.Printf("%T %v %t\n", t, t, t) // bool型じゃないと正しく表示しない%t
    fmt.Printf("%T %v %t\n", f, f, 0) //  %!t(int=0)となりbool型じゃないと出る

}

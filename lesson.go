package main

import (
	"fmt"
)

type Human interface {
	Say() string
}
type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string { // 実装した中身を操作するのでアスタリスク
	p.Name = "Mr." + p.Name
	return p.Name
}
func DriveCar (human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
}
func main() {
	var mike Human = &Person{"Mike"}
	var dog Dog = Dog{"wanwan"}
	DriveCar(mike)
	DriveCar(dog)
}


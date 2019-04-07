package main

import "fmt"

type Monkey struct {
	Name string
}
type BirdAble interface {
	Flying()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}

type LittleMonkey struct {
	Monkey
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习，学会飞翔")
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	var monkey1 BirdAble
	monkey1 = &monkey
	monkey1.Flying()
}

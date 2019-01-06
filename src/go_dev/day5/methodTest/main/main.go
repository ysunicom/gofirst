package main

import "fmt"

type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Println("running")
}
func (p *Car) init(weight int, name string) {
	p.weight = weight
	p.name = name
	fmt.Println(p)
}
func (p *Car) set(weight int, name string) {
	p.weight = weight
	p.name = name
}

type Bike struct {
	Car
}

type Train struct {
	c Car
}

func main() {
	var a Bike
	a.init(100, "bike")
	a.Run()
	var b Train
	b.c.set(1000, "Train")
	fmt.Println(b)
	var c Bike
	c.set(200, "bike1")
	fmt.Println(c)
	c.Run()
}

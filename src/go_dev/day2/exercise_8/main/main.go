package main

import "fmt"

func modify(a *int) {
	*a = 10
	fmt.Println(a)
}
func main() {
	a := 5
	b := make(chan int, 1)
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	modify(&a)
	fmt.Println(&a)
	fmt.Println("a =", a)
}

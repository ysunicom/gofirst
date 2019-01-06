package main

import "fmt"

func swapab(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	var (
		a, b int
	)
	a = 6
	b = 7
	swapab(&a, &b)
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

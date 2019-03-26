package main

import "fmt"

type mySum func(n1, n2 int) int

func sum(n1, n2 int) int {
	return n1 + n2
}

func myFunc(a mySum, n1, n2 int) int {
	return a(n1, n2)
}

func main() {
	a := sum
	fmt.Println(myFunc(a, 1, 2))
}

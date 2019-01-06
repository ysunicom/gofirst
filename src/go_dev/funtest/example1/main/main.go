package main

import "fmt"

func myfun(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
func myfun1(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
func main() {
	myfun(1, 31, 44)
	myfun1(1.3, "hello", 47)

	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(23, 44))
}

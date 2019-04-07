package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println("n = ", n)
	}

}

func test2(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n = ", n)

}

func main() {
	test(5592327)
}

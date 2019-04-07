package main

import "fmt"

func fn(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*fn(n-1) + 1
	}

}

func main() {
	for i := 1; i <= 40; i++ {
		fmt.Printf("%d`s fab is %d\n", i, fn(i))
	}
}

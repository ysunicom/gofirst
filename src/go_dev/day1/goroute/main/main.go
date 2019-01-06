package main

import (
	"fmt"
	"go_dev/day1/goroute/calc"
)

// func add(a int, b int, c chan int) {
// 	sum := a + b
// 	c <- sum
// }

func main() {
	var sum int
	for i := 0; i < 100; i++ {
		for j := 0; j < 20; j++ {
			sum = calc.Add(i, j)
			fmt.Println("sum =", sum)
		}
	}

}

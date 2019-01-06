package main

import (
	"fmt"
	"go_dev/day1/goroute_example/goroute"
)

func main() {
	pipe := make(chan int, 1)
	var sum int
	go goroute.Add(3, 4, pipe)
	sum = <-pipe
	fmt.Println("sum =", sum)

}

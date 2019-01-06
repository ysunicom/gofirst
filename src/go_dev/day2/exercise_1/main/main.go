package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d", &input)
	for i := 0; i <= input; i++ {
		fmt.Printf("%d + %d = %d\n", i, input-i, input)
	}
	fmt.Printf("%x %d %b %c", input, input, input, input)
}

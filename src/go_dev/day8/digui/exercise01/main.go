package main

import "fmt"

func fab(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fab(n-1) + fab(n-2)
	}

}

func main() {
	for i := 1; i <= 40; i++ {
		fmt.Printf("%d`s fab is %d\n", i, fab(i))
	}
}

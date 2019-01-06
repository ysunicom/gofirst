package main

import (
	"fmt"
	"sort"
)

func testSort() {
	var a = [...]int{1, 2, 9, 4, 5, 6}
	fmt.Printf("%p\n", &a)
	sort.Ints(a[:])
	fmt.Printf("%p\n", &a)
	fmt.Println(a)
}

func main() {
	testSort()
}

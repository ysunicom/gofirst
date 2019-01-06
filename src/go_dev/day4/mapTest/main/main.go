package main

import (
	"fmt"
	"sort"
)

func mapSort() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[8] = 8
	a[3] = 3
	a[2] = 2
	a[1] = 1
	a[9] = 9
	var keys []int
	for k, v := range a {
		keys = append(keys, k)
		fmt.Println(k, v)
	}
	fmt.Println("-----------------")
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	}

}

func main() {
	mapSort()
}

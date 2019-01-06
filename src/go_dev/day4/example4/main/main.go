package main

import "fmt"

func slicetest() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Printf("%p\n", b)
	fmt.Printf("%p", &a[1])
}

func slicetest1() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	b := a[1:2]
	fmt.Println(b)
	b = append(b, 7)
	//b = append(b, 7)
	b = b[0:3]
	fmt.Println(b)
	fmt.Printf("len[%d] cap[%d]", len(b), cap(b))
}
func slicetest2() {
	var a []int = []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 10)
	copy(b, a)
	fmt.Printf("b:%p,a:%p\n", b, a)
	fmt.Println(b)
}
func main() {
	slicetest2()
}

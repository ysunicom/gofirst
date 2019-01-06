package main

import "fmt"

type Student struct {
	Name string
	Sex  string
}

func Test(a interface{}) {
	b, ok := a.(Student)
	if ok == false {
		fmt.Println("convert fail")
		return
	}
	// b += 3
	fmt.Println(b)
}

func main() {
	var a Student
	Test(a)
}

package main

import "fmt"

type Student struct {
	Name  string
	Age   int64
	score float32
}

func main() {

	var stu Student
	stu.Name = "ys中国人是一个强大的民族"
	stu.Age = 18
	stu.score = 81.2

	fmt.Println(stu)
	fmt.Printf("Name:%p\n", &stu.Name)
	fmt.Println(&stu.Age)
	fmt.Println(&stu.score)

}

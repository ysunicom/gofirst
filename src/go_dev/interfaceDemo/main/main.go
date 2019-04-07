package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

type Students []Student

func (stus Students) Len() int {
	return len(stus)
}

func (stus Students) Less(i, j int) bool {
	//return stus[i].Score > stus[j].Score
	return stus[i].Age < stus[j].Age
}

func (stus Students) Swap(i, j int) {
	stus[i], stus[j] = stus[j], stus[i]
}

func main() {
	var stus Students
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("学生~%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		stus = append(stus, stu)
	}
	for _, v := range stus {
		fmt.Println(v)
	}
	fmt.Println("-------------------------")
	sort.Sort(&stus)
	for _, v := range stus {
		fmt.Println(v)
	}
}

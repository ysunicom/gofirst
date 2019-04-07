package main

import "fmt"

func printBlank(totalrow, row int) {
	for i := 0; i < totalrow-row; i++ {
		fmt.Print(" ")
	}
}

func printStar(row int) {
	for i := 0; i < 2*row-1; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	fmt.Println("请输入金字塔的层数")
	var totalrow int
	fmt.Scanln(&totalrow)
	for i := 1; i <= totalrow; i++ {
		printBlank(totalrow, i)
		printStar(i)
	}
}

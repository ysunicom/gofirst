package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filepath := "abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err:%v", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "hello Gardon\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}

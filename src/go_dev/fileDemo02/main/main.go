package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("e:/BGP.txt")
	if err != nil {
		fmt.Println("file open err:%v", err)
	}
	fmt.Printf("%v\n", string(content))
}

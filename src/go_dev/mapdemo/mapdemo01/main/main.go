package main

import "fmt"

func main() {
	var city = make(map[string]string, 2)

	city["no2"] = "北京"
	city["no3"] = "西安"
	city["no4"] = "武汉"
	city["no1"] = "宜昌"

	fmt.Println(city)
}

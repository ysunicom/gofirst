package main

import "fmt"

func main() {
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}
	newMonster := map[string]string{
		"name": "狐狸精",
		"age":  "300",
	}
	monsters = append(monsters, newMonster)

	for i, v := range monsters {
		fmt.Printf("i = %v,v = %v\n", i, v)
		for k, v := range v {
			fmt.Printf("\tk = %v,v = %v\n", k, v)
		}
		fmt.Println()
	}
}

package main

import "fmt"

//有一堆桃子，猴子每天吃其中的一半，并再多吃一个，第10天，只剩一个桃了，问有多少桃
/*
第10天 1个
第9天 是（第10天 + 1）* 2
第n天是(第n-1天　+ 1) * 2
(peach(n + 1) + 1) * 2
*/
func peach(n int) int {
	if n == 10 {
		return 1
	} else {
		return 2 * (peach(n+1) + 1)
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("第%d天的桃子数是:%d\n", i, peach(i))
	}

}

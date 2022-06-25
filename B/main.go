package main

import (
	"fmt"
)

func main() {

	var count int
	fmt.Scanln(&count)
	for i := 0; i < count; i++ {
		var cnt int
		fmt.Scanln(&cnt)
		prices := make([]int, cnt)
		for j := 0; j < cnt; j++ {
			fmt.Scan(&prices[j])
			//prices[j] = j
			fmt.Println("j = ", j)
		}
		fmt.Println(cnt, prices)
	}
}

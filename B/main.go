package main

import (
	"fmt"
)

func main() {

	var count int
	var cnt int
	fmt.Scanln(&count)
	for i := 0; i < count; i++ {
		fmt.Scanln(&cnt)
		fmt.Println("cnt =", cnt)
		prices := make([]int, cnt)
		for j := 0; j < cnt; j++ {
			fmt.Scan(&prices[j])
			fmt.Println("j = ", j)
		}
		fmt.Println(cnt, prices)
	}
}

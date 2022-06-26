package main

import (
	"fmt"
	"sort"
)

func main() {

	//var count int

	//fmt.Scanln(&count)

	//for i := 0; i < count; i++ {
	var cnt int
	fmt.Scanln(&cnt)
	prices := make([]int, cnt)
	for j := 0; j < cnt; j++ {
		fmt.Scan(&prices[j])
	}
	fmt.Println(prices)
	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})
	fmt.Println(prices)
	var discount int
	var itemCount int
	thing := prices[0]
	for k, val := range prices {
		if val == thing {
			itemCount++
			discount += val
		} else {
			discount -= val * itemCount / 3
			fmt.Println(discount)
			discount = val
			itemCount = 1
			thing = val
		}
		if k == len(prices) {
			discount -= val * itemCount / 3
			fmt.Println(discount)
		}
	}
	//}

}

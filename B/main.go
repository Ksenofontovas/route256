package main

import (
	"fmt"
	"sort"
)

func main() {
	var count int

	fmt.Scanln(&count)

	for i := 0; i < count; i++ {
		var cnt int
		var sum int
		fmt.Scanln(&cnt)
		prices := make([]int, cnt)
		for j := 0; j < cnt; j++ {
			fmt.Scan(&prices[j])
		}
		sort.Slice(prices, func(i, j int) bool {
			return prices[i] < prices[j]
		})
		var discount int
		var itemCount int
		thing := prices[0]
		var prevVal int
		for k, val := range prices {
			if val == thing {
				prevVal = val
				itemCount++
				discount += val
			} else {
				if itemCount >= 3 {
					discount -= prevVal * (itemCount / 3)
					sum += discount
				} else {
					sum += discount
				}
				discount = val
				itemCount = 1
				thing = val
			}
			if k == len(prices)-1 && val == thing {
				if itemCount >= 3 {
					discount -= val * (itemCount / 3)
					sum += discount
				} else {
					sum += discount
				}
			}
		}
		fmt.Println(sum)
	}

}

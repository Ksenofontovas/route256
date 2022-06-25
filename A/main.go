package main

import (
	"fmt"
)

func main() {

	var count, f, s int
	fmt.Scanln(&count)
	for i := 0; i < count; i++ {
		fmt.Scanln(&f, &s)
		fmt.Println(f + s)

	}
}

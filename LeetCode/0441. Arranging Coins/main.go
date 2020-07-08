package main

import (
	"fmt"
)

func main() {
	ans := arrangeCoins(10)
	fmt.Println(ans)
}

func arrangeCoins(n int) int {
	coin := 0
	for n >= 0 {
		coin++
		n -= coin
	}
	return coin - 1
}

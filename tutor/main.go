package main

import "fmt"

func main() {
	result := cal("+", 5, 7)
	fmt.Println(result)
}

func cal(op string, nums ...int) int {
	if op == "+" {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return sum
	}
	return 0
}

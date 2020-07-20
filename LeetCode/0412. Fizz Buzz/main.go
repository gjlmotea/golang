package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(5))
}

func fizzBuzz(n int) []string {
	str := make([]string, 0)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			str = append(str, "FizzBuzz")
		case i%5 == 0:
			str = append(str, "Buzz")
		case i%3 == 0:
			str = append(str, "Fizz")
		default:
			str = append(str, strconv.Itoa(i))
		}
		fmt.Print(1 != 2)

	}
	return str
}

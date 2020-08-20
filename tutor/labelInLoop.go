package main

import (
	"fmt"
)

func main() {
	FirstNames := []string{"aaa", "bbb", "ccc"}
	LastNames := []string{"111", "222", "333"}

Loop:
	for _, firstName := range FirstNames {
		//Loop:
		for _, lastName := range LastNames {
			fmt.Printf("Name: %s %s\n", firstName, lastName)

			if firstName == "bbb" && lastName == "111" {
				break Loop
				//continue Loop
			}
		}
	}
	println("Over.")
}


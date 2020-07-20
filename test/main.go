package main

import "fmt"

func c() (j int) {
	i:=0
	defer fmt.Println("i",i)
	defer func() { i++ }()
	defer fmt.Println("i",i)
	return 1

}
func main() {
	fmt.Println(c())
	return
}
package main

import "fmt"

func main() {
	var b_min byte = 0
	var b_max byte = 255
	var i_min int8 = -128
	var i_max int8 = 127
	fmt.Println(b_min, b_max)
	fmt.Println(i_min, i_max)
	// 0 255
	// -128 127

	// ^: `Not` in 1 symbol
	var a byte = 1
	var b byte = 0
	var c int8 = 1
	var d int8 = 0
	fmt.Println("a:", a, ^a, 123&^a)
	fmt.Println("b:", b, ^b, 123&^b)
	fmt.Println("c:", c, ^c, 123&^c)
	fmt.Println("d:", d, ^d, 123&^d)
	// a: 1 254 122
	// b: 0 255 123
	// c: 1 -2 122
	// d: 0 -1 123

	//^: `XOR` in 2 symbols
	fmt.Println(1^100, 1^101)

	var o int = 011   // 8進位
	var p byte = 0x12 // 16進位
	fmt.Println(o)
	fmt.Println(p)
}

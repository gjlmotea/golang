package main

import (
	"fmt"
	"math"
)

func main() {
	listMultSqrts(20)
	listExps(8, 3)
	listExps(2, 4)
}

func listMultSqrts(number int)  {
	//列出 數字內 所有的開n次平方根
	for i := 1; i <= number; i++{
		var multSqrt = float64(i)
		for j := i; j > 0; j--{
			fmt.Printf("%.3f \t", multSqrt)
			multSqrt = math.Sqrt(float64(multSqrt))
		}
		fmt.Println()
	}
}

func listExps(ceil int, parts int){
	//列出 數字內 所有1以內的分數次方
	//ceil 的 part/parts 次方
	//	ex:	exps(2, 3)
	//	result:	2^(0/3), 2^(1/3), 2^(2/3), 2^(3/3)
	f64_ceil := float64(ceil)
	f64_parts := float64(parts)

	for part := 0; part <= parts; part++{
		f64_part := float64(part)
		var f64_frac float64 = f64_part / f64_parts
		fmt.Printf("%.3f \t", math.Pow(f64_ceil, f64_frac))
	}
}

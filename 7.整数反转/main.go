package main

import (
	"fmt"
	"math"
)

func numRe(num int32) int32 {
	var res int32 = 0
	for num != 0 {
		res = res * 10 + num % 10
		num = num / 10
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}

func main() {
	fmt.Println(numRe(-120))
}

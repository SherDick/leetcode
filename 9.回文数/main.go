package main

import (
	"fmt"
	"math"
)

/**
	反转数字，然后比较
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var y int = 0
	var t int = x
	for t != 0 {
		y = y * 10 + t % 10
		t = t /10
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		y = 0
	}
	if x != y {
		return false
	}
	return true
}

/*
	反转一半数字，然后比较
	时间复杂度：O(logn)
	空间复杂度：O(1)
*/
func isPalindrome2(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || (x % 10) == 0{
		return false
	}

	var t int = x
	var y int = 0
	for {
		y = y * 10 + t % 10
		t = t / 10
		if t <= y {
			break
		}
	}
	if y == t || (y / 10) == t {
		return true
	}
	return false
}

func main() {
	//fmt.Println(isPalindrome(3212123))
	fmt.Println(isPalindrome2(0))
}

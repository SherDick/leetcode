package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	m := map[byte]int{}
	left, right := 0, -1
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right+1 < len(s) {
			if _, ok := m[s[right+1]]; ok {
				break
			}
			m[s[right+1]]++
			right++
		}
		left = max(left, right+1-i)
	}
	return left
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

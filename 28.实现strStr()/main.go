package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if 0 == len(needle) {
		return 0
	}
	curLen := 0
	p, q := haystack, needle
	for i := 0; i < len(haystack); i++ {
		if p[i] == q[0] {
			i = i + 1
			for j := 1; j < len(needle) && i < len(haystack); j++ {
				if p[i] != q[j] {
					// 回退至p[i] == q[0]中的i+1处
					i = i - j
					break
				} else {
					curLen = j
					i = i + 1
				}
			}
			if curLen+1 == len(needle) {
				return i - (curLen + 1)
			}
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	//s1, s2 := "mississippi", "issip"
	//s1, s2 := "hello", "ll"
	//s1, s2 := "aaa", "aaaa"
	s1, s2 := "mississippi", "issipi"
	fmt.Println(strStr(s1, s2))
}

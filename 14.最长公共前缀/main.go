package main

import "fmt"

/*
	纵向扫描
	时间复杂度：O(mn)
	空间复杂度：O(1)
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

/*
	二分查找
	时间复杂度：O(mnlogm), m: 最短字符串长度, n: 数量
	空间复杂度：O(1)
	"a", "b"
	"a"
	"c", "c"
	"dog", "racecar", "car"
*/
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	isComPrefix := func(length int) bool {
		if length == 0 {
			return false
		}
		s, n := strs[0][:length], len(strs)
		for i := 1; i < n; i++ {
			if strs[i][:length] != s {
				return false
			}
		}
		return true
	}
	var minLen int = len(strs[0])
	for _, s := range strs {
		if minLen > len(s) {
			minLen = len(s)
		}
	}
	start, end := 0, minLen
	for start < end {
		mid := (start + end + 1) / 2
		if isComPrefix(mid) {
			// true
			start = mid
		} else {
			end = mid - 1
		}
	}
	return strs[0][:start]
}

func main() {
	var strs []string
	strs = append(strs, "dog")
	strs = append(strs, "racecar")
	strs = append(strs, "car")
	//fmt.Println(longestCommonPrefix(strs))
	fmt.Println(longestCommonPrefix2(strs))
}

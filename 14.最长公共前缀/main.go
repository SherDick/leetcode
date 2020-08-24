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
	for i := 0;i < len(strs[0]);i++ {
		for j := 1; j < len(strs);j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

/*
	二分查找
	时间复杂度：
	空间复杂度：
 */
func longestCommonPrefix2(strs []string) string {
	return ""
}

func main() {
	var strs []string
	strs = append(strs, "flower")
	strs = append(strs, "flow")
	strs = append(strs, "flight")
	fmt.Println(longestCommonPrefix(strs))
}

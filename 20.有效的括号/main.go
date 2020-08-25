package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	// 长度为奇数
	if len(s)%2 != 0 {
		return false
	}
	brackets := make(map[string]string, 3)
	brackets["("] = ")"
	brackets["{"] = "}"
	brackets["["] = "]"
	// 最后一个为左括号
	char := string(s[len(s)-1])
	if _, ok := brackets[char]; ok {
		return false
	}
	// 第一个为右括号
	char = string(s[0])
	if ")" == char || "}" == char || "]" == char {
		return false
	}

	stack := make([]string, 0)
	push := func(s string) {
		stack = append(stack, s)
	}
	pop := func() string {
		s := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		return s
	}

	for i := 0; i < len(s); i++ {
		char = string(s[i])
		if _, ok := brackets[char]; ok {
			push(char)
		} else {
			t := pop()
			if t2, ok := brackets[t]; ok {
				if t2 != char {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	s := "{[]}"
	fmt.Println(isValid(s))
}

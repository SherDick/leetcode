package main

import "fmt"

func plusOne(digits []int) []int {
	stack1, stack2 := make([]int, 0), make([]int, 0)
	push := func(val int) {
		stack2 = append(stack2, val)
	}
	pop := func() int {
		val := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		return val
	}

	stack1 = digits
	flag := 1
	for i := 0; i < len(digits); i++ {
		num := pop()
		num = num + flag
		if 10 == num {
			push(0)
			if i+1 == len(digits) {
				push(1)
				break
			}
			flag = 1
		} else {
			push(num)
			flag = 0
		}
	}

	return reverse(stack2)
}

func reverse(array []int) []int {
	for i := 0; i < (len(array)+1)/2; i++ {
		high := len(array) - i - 1
		array[i], array[high] = array[high], array[i]
	}
	return array
}

func main() {
	num := []int{2, 3, 4, 9, 5, 8, 9, 9}
	fmt.Println(plusOne(num))
}

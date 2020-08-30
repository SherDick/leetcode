package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for j = 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i = i + 1
		}
	}
	return i + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//nums := []int{1, 1, 2}
	count := removeDuplicates(nums)
	for i := 0; i < count; i++ {
		fmt.Println(nums[i])
	}
}

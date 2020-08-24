package main

import "fmt"

func twoSum(nums []int, target int) []int  {
	numMap := make(map[int]int, 0)
	for index, num := range nums {
		if i, ok := numMap[target - num]; ok {
			return []int{i, index}
		}
		numMap[num] = index
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}

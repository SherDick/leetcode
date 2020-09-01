package main

import "fmt"

/**
0,1,2,2,3,0,4,2   val = 2
*/
func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if val != nums[i] {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	return j
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	count := removeElement(nums, 2)
	for i := 0; i < count; i++ {
		fmt.Println(nums[i])
	}
}

package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	index := m + n - 1
	i, j := m-1, n-1
	for j >= 0 && i >= 0 {
		if nums1[i] <= nums2[j] {
			nums1[index] = nums2[j]
			j--
		} else {
			nums1[index] = nums1[i]
			i--
		}
		index--
	}
	for ; j >= 0; j-- {
		nums1[index] = nums2[j]
		index--
	}
}

func main() {
	//num1 := []int{1, 2, 3, 0, 0, 0}
	//num2 := []int{2, 5, 6}
	//num1 := []int{0}
	//num2 := []int{1}
	num1 := []int{2, 0}
	num2 := []int{1}
	fmt.Println(num1)
	//merge(num1, 3, num2, 3)
	//merge(num1, 0, num2, 1)
	merge(num1, 1, num2, 1)
	fmt.Println(num1)

}

package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/
func smallerNumbersThanCurrent(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		var count = 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}

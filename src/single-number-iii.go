package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/single-number-iii/
func singleNumber(nums []int) []int {
	var result []int
	var m = make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] {
			delete(m, nums[i])
		} else {
			m[nums[i]] = true
		}
	}
	for k, _ := range m {
		if m[k] {
			result = append(result, k)
		}
	}
	return result
}

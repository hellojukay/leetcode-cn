package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//https://leetcode-cn.com/problems/element-appearing-more-than-25-in-sorted-array/
func findSpecialInteger(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	var max_count = 0
	var length = len(arr)
	var count = 1
	var max = arr[0]
	for i := 1; i < length; i++ {
		if arr[i-1] == arr[i] {
			count++
		} else {
			count = 1
		}
		if count > max_count {
			max_count = count
			max = arr[i-1]
		}
	}
	return max
}

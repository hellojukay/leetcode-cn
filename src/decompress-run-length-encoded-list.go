package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func decompressRLElist(nums []int) []int {
	var length = len(nums)
	var result []int
	for i := 0; i*2 <= length-1; i++ {
		result = push(result, nums[i*2], nums[i*2+1])
	}
	return result
}

func push(src []int, m int, n int) []int {
	for i := 0; i < m; i++ {
		src = append(src, n)
	}
	return src
}

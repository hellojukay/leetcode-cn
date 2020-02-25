package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func findRepeatNumber(nums []int) int {
	var size = len(nums)
	var arr = make([]bool, size)
	for _, n := range nums {
		if arr[n] {
			return n
		}
		arr[n] = true
	}
	return 0
}

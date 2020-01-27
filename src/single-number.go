package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func singleNumber(nums []int) int {
	var result = 0
	for _, n := range nums {
		result = result ^ n
	}
	return result
}

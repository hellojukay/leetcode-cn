package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func plusOne(digits []int) []int {
	var c = 1
	for index := len(digits) - 1; index >= 0; index-- {
		var tmp = (digits[index] + c) / 10
		digits[index] = (digits[index] + c) % 10
		c = tmp
	}
	if c == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

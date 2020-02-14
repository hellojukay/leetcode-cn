package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//https://leetcode-cn.com/problems/swap-numbers-lcci/
func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] ^ numbers[1]
	numbers[1] = numbers[1] ^ numbers[0]
	numbers[0] = numbers[0] ^ numbers[1]
	return numbers
}

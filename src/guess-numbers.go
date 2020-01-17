package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/guess-numbers/
func game(guess []int, answer []int) int {
	var result = 0
	for i := 0; i < 3; i++ {
		if guess[i] == answer[i] {
			result++
		}
	}
	return result
}

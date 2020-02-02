package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/nim-game/
func canWinNim(n int) bool {
	return n%4 != 0
}

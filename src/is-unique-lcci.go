package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/is-unique-lcci/
func isUnique(astr string) bool {
	var m = make(map[rune]bool)
	for _, ch := range astr {
		if m[ch] {
			return false
		}
		m[ch] = true
	}
	return true
}

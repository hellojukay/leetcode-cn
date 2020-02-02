package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/jewels-and-stones/
func numJewelsInStones(J string, S string) int {
	if len(J) == 0 {
		return 0
	}
	var m = make(map[rune]bool)
	for _, ch := range J {
		m[ch] = true
	}
	var count = 0
	for _, ch := range S {
		ok, _ := m[ch]
		if ok {
			count++
		}
	}
	return count
}

package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid/
func minAddToMakeValid(S string) int {
	var stack = make([]rune, len(S))
	var top = -1
	for _, ch := range S {
		if top < 0 {
			top++
			stack[top] = ch
		} else {
			if ch == ')' && stack[top] == '(' {
				top--
			} else {
				top++
				stack[top] = ch
			}
		}
	}
	return top + 1
}

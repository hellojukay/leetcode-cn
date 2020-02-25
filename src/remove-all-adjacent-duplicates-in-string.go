package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/
func removeDuplicates(S string) string {
	if len(S) == 0 {
		return S
	}
	var stack []byte
	var arr = []byte(S)
	for _, ch := range arr {
		if len(stack) == 0 {
			stack = append(stack, ch)
			continue
		}
		if ch == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return string(stack)
}

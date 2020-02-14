package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	var result = ""
	for _, ch := range s {
		if ch == ' ' {
			result = result + "%20"
		} else {
			result = result + string(ch)
		}
	}
	return result
}

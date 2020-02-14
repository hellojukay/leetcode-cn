package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func reverseLeftWords(s string, n int) string {
	if n == 0 {
		return s
	}
	var left = s[0:n]
	var right = s[n:]
	return right + left
}

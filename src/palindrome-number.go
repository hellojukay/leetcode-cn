package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var str = fmt.Sprintf("%d", x)
	var reverse = ""
	for _, ch := range str {
		reverse = string(ch) + reverse
	}
	return reverse == str
}

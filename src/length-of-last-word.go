package main

import (
	"strings"
)

// https://leetcode-cn.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	var arr = strings.Split(s, " ")
	var word = arr[len(arr)-1]
	return len(word)
}

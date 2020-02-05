package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/reverse-string/
func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}
	var i = 0
	var j = len(s) - 1
	for {
		if i >= j {
			break
		}
		var tmp = s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}
}

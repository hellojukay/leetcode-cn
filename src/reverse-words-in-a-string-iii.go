package main

import "strings"

func reverseWords(s string) string {
	var strs = strings.Split(s, " ")
	var result = ""
	for _, str := range strs {
		result = result + " " + reverse(str)
	}
	return strings.TrimLeft(result, " ")
}
func reverse(str string) string {
	var result = ""
	for _, ch := range str {
		result = string(ch) + result
	}
	return result
}

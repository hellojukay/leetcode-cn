package main

import "strings"

// https://leetcode-cn.com/problems/c
func printVertically(s string) []string {
	if len(s) == 0 {
		return nil
	}
	strs := strings.Split(s, " ")
	var result []string
	var tmp []string
	var word string
	var flag = false
	for {
		if flag {
			break
		}
		flag = true
		for _, str := range strs {
			if len(str) > 1 {
				flag = false
			}
			if len(str) > 0 {
				word = word + string(str[0])
				tmp = append(tmp, string(str[1:len(str)]))

			} else {
				word = word + " "
				tmp = append(tmp, "")
			}
		}
		word = strings.TrimRight(word, " ")
		//word = strings.TrimLeft(word, " ")
		if len(word) > 0 {
			result = append(result, word)
		}
		word = ""
		strs = tmp
		tmp = nil
	}
	return result
}

func main() {
	printVertically("HOW ARE YOU")
}

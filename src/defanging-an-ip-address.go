package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/defanging-an-ip-address/
func defangIPaddr(address string) string {
	var result string
	for _, ch := range address {
		if ch == '.' {
			result = result + "[.]"
		} else {
			result = result + string(ch)
		}
	}
	return result
}

package main

import "unicode"

//https://leetcode-cn.com/problems/valid-palindrome/
func isPalindrome(s string) bool {
	var str = ""
	var result = ""
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			str = str + string(ch)
			continue
		}
		if unicode.IsLower(ch) {
			str = str + string(ch)
			continue
		}
		if unicode.IsUpper(ch) {
			str = str + string(unicode.ToLower(ch))
		}
	}
	for _, ch := range str {
		result = string(ch) + result
	}
	println(str)
	println(result)
	return result == str
}

func main() {
	println(isPalindrome("A man, a plan, a canal: Panama"))
}

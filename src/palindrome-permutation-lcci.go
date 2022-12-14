package main

// https://leetcode.cn/problems/palindrome-permutation-lcci/
func canPermutePalindrome(s string) bool {
	var hash = make(map[rune]int)
	for _, ch := range s {
		hash[ch]++
	}
	var count int
	for _, v := range hash {
		if v%2 != 0 {
			count++
		}
	}
	return count <= 1
}

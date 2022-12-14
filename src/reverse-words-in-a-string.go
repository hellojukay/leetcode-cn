package main

import (
	"regexp"
	"strings"
)

// https://leetcode.cn/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	var reverse = func(s []string) []string {
		a := make([]string, len(s))
		copy(a, s)

		for i := len(a)/2 - 1; i >= 0; i-- {
			opp := len(a) - 1 - i
			a[i], a[opp] = a[opp], a[i]
		}

		return a
	}
	re := regexp.MustCompile(`\s+`)
	var arr = re.Split(s, -1)
	arr = reverse(arr)
	return strings.TrimSpace(strings.Join(arr, " "))
}

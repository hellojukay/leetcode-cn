package main

import "strings"

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

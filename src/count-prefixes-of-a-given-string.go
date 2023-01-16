package main

import "strings"

func countPrefixes(words []string, s string) int {
	var n int
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			n++
		}
	}
	return n
}

package main

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	var arr = strings.Split(s, " ")
	if len(arr) != len(pattern) {
		return false
	}
	var hash = make(map[rune]string)
	var hash2 = make(map[string]rune)
	for i := 0; i < len(pattern); i++ {
		var ch = ([]rune(pattern))[i]
		v, ok := hash[ch]
		if ok {
			if arr[i] != v {
				return false
			}
		} else {
			hash[ch] = arr[i]
		}
		v1, ok1 := hash2[arr[i]]
		if ok1 {
			if v1 != ch {
				return false
			}
		} else {
			hash2[arr[i]] = ch
		}
	}
	return true
}

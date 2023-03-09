package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	loop("", []rune(s))
	return result
}

var result []string

func loop(prefix string, last []rune) {
	if len(last) == 0 && strings.Count(prefix, "") == 3 {
		result = append(result, prefix)
		return
	}
	if len(last) == 0 {
		return
	}
	for i := 1; i <= 3; i++ {
		if len(last) > i {
			current := last[:i]
			if validate(string(current)) {
				if prefix == "" {
					prefix = string(current)
				} else {
					prefix = prefix + "." + string(current)
				}
				loop(prefix, last[i:])
			}
		}
	}
}
func validate(s string) bool {
	if len([]rune(s)) > 1 && strings.HasPrefix(s, "0") {
		return false
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	return n >= 0 && n <= 255
}

package main

import (
	"sort"
	"unicode"
)

func secondHighest(s string) int {
	var numbers sort.IntSlice
	for _, ch := range s {
		if unicode.IsNumber(ch) {
			numbers = append(numbers, int(ch-rune('0')))
		}
	}
	if len(numbers) <= 1 {
		return -1
	}
	numbers.Sort()
	end := numbers[len(numbers)-1]
	for i := len(numbers) - 2; i >= 0; i-- {
		if numbers[i] != end {
			return numbers[i]
		}
	}
	return -1
}

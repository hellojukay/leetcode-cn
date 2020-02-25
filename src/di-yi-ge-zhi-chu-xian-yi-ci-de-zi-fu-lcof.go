package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	var result byte = ' '
	var arr = []byte(s)
	var m = make(map[byte]int)
	for _, ch := range arr {
		m[ch]++
	}
	for _, ch := range arr {
		if v, ok := m[ch]; ok {
			if v == 1 {
				return ch
			}
		}
	}
	return result
}

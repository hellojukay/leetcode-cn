package main

import (
	"bytes"
	"strings"
)

func minBitFlips(start int, goal int) int {
	var s1 = toBitString(start)
	var s2 = toBitString(goal)
	println(s1, s2)
	var c int
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			c = c + 1
		}
	}
	return c
}

func toBitString(n int) []byte {
	var buffer bytes.Buffer
	var c = 0
	for n > 0 {
		c = n % 2
		if c == 0 {
			buffer.WriteString("0")
		} else {
			buffer.WriteString("1")
		}
		n = n / 2
	}
	var length = buffer.Len()
	if length < 32 {
		buffer.WriteString(strings.Repeat("0", 32-length))
	}
	return buffer.Bytes()
}

func main() {
	var n = minBitFlips(10, 82)
	println(n)
}

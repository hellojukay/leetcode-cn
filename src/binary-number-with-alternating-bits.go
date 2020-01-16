package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/binary-number-with-alternating-bits/
func hasAlternatingBits(n int) bool {
	if n == 0 {
		return true
	}
	for {
		if n == 0 {
			break
		}
		var hight_bit = n >> 1
		if low_bit(hight_bit) == n&1 {
			return false
		}
		n = hight_bit
	}
	return true
}

// 获取一个二进制数最低位置数
func low_bit(n int) int {
	return n & 1
}

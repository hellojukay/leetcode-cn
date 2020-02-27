package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var result = generateParenthesis(3)
	for _, str := range result {
		if validate(str) {
			println(str)
		}
	}
}

// https://leetcode-cn.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	var max = pow2(n*2) + 1
	var size = n * 2
	var result []string
	for i := 0; i <= max; i++ {
		var s = num2str(size, i)
		if validate(s) {
			result = append(result, s)
		}
	}
	return result
}
func pow2(n int) int {
	var x = 1
	for i := 0; i < n; i++ {
		x = x * 2
	}
	return x
}

// 生成括号
func num2str(size int, n int) string {
	var result string
	// （　为 0 , ) 为 1
	var x = n
	for i := 1; i <= size; i++ {
		var c = x % 2
		if c == 0 {
			result = "(" + result
		} else {
			result = ")" + result
		}
		x = x / 2
	}
	return result
}

func validate(str string) bool {
	var stack []rune
	for _, ch := range str {
		if len(stack) == 0 {
			stack = append(stack, ch)
			continue
		}
		if ch == rune(')') && stack[len(stack)-1] == rune('(') {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}

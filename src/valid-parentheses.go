package main

// https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	var stack = make([]rune, len(s))
	var top = -1
	for _, ch := range s {
		if top == -1 {
			top = 0
			stack[top] = ch
		} else {
			if stack[top] == '(' && ch == ')' {
				top--
				continue
			}
			if stack[top] == '[' && ch == ']' {
				top--
				continue
			}
			if stack[top] == '{' && ch == '}' {
				top--
				continue
			}
			top++
			stack[top] = ch
		}
	}
	return top == -1
}

func main() {
	println(isValid("()()[]{}"))
	println(isValid("([)"))
}

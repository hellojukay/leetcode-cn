package main

import "fmt"
import "strconv"

func main() {
	fmt.Println("vim-go")
}

func eval(f func(int64, int64) int64, a int64, b int64) int64 {
	return f(a, b)
}
func sum(a int64, b int64) int64 {
	return a + b
}
func div(a int64, b int64) int64 {
	return a - b
}

func mod(a int64, b int64) int64 {
	return a / b
}
func mult(a int64, b int64) int64 {
	return a * b
}
func evalRPN(tokens []string) int {
	var stack []int64
	var x, y int64
	var tmp int64
	for _, str := range tokens {
		if str == "-" {
			x = stack[len(stack)-2]
			y = stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			tmp = eval(div, x, y)
			stack = append(stack, tmp)
			continue
		}
		if str == "+" {
			x = stack[len(stack)-2]
			y = stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			tmp = eval(sum, x, y)
			stack = append(stack, tmp)
			continue
		}
		if str == "*" {
			x = stack[len(stack)-2]
			y = stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			tmp = eval(mult, x, y)
			stack = append(stack, tmp)
			continue
		}
		if str == "/" {
			x = stack[len(stack)-2]
			y = stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			tmp = eval(mod, x, y)
			stack = append(stack, tmp)
			continue
		}
		x, _ := strconv.ParseInt(str, 10, 64)
		stack = append(stack, x)
	}
	tmp = stack[0]
	return int(tmp)
}

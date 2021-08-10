package main

func tribonacci(n int) int {
	return fib(0, 1, 1, n, 0)
}

func fib(a, b, c, n, account int) int {
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	if n == 2 {
		return c
	}
	tmp := a + b + c
	return fib(b, c, tmp, n-1, tmp)
}

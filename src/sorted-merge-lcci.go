package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/sorted-merge-lcci/
func merge(A []int, m int, B []int, n int) {
	var index = m + n - 1
	for {
		if m == 0 && n == 0 {
			break
		}
		// get max
		if m == 0 {
			A[index] = B[n-1]
			n--
			index--
			continue
		}
		if n == 0 {
			A[index] = A[m-1]
			m--
			index--
			continue
		}
		if A[m-1] < B[n-1] {
			A[index] = B[n-1]
			index--
			n--
		} else {
			A[index] = A[m-1]
			index--
			m--
		}
	}
}

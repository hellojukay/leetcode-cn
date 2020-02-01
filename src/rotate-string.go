package main

func rotateString(A string, B string) bool {
	if A == B {
		return true
	}
	for i := 1; i < len(A)-1; i++ {
		if string(A[i:])+string(A[:i]) == B {
			return true
		}
	}
	return false
}

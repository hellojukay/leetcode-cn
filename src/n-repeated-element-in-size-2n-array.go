package main

func repeatedNTimes(A []int) int {
	var n = len(A) / 2
	var hash = make(map[int]int)
	for _, x := range A {
		hash[x]++
	}
	for k, v := range hash {
		if v == n {
			return k
		}
	}
	return 0
}

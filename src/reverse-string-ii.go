package main

func reverseStr(s string, k int) string {
	if k == 1 {
		return s
	}
	var array = []rune(s)
	if len(array) < k {
		return reverse(array)
	}
	if len(array) < k*2 {
		return reverse(array[:k]) + string(array[k:])
	}
	return reverse(array[:k]) + string(array[k:2*k]) + reverseStr(string(array[2*k:]), k)
}

func reverse(array []rune) string {
	var result []rune = make([]rune, len(array))
	for _, ch := range array {
		result = append([]rune{ch}, result...)
	}
	return string(result)
}

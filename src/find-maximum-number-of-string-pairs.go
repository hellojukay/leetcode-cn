func maximumNumberOfStringPairs(words []string) int {
	var count int
	for i := 0; i < len(words); i++ {
		for j := len(words) - 1; j > i; j-- {
			if words[i] == reverse(words[j]) {
				count++
			}
		}

	}
	return count
}
func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
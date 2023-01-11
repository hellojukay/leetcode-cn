func prefixCount(words []string, pref string) int {
	var n = 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			n++
		}
	}
	return n
}
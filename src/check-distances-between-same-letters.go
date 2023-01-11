func checkDistances(s string, distance []int) bool {
	var hash = make(map[rune]int)
	for index, ch := range []rune(s) {
		n, ok := hash[ch]
		if ok {
			hash[ch] = abs(index-n) - 1
		} else {
			hash[ch] = index
		}
	}
	for ch, d := range hash {
		var n = int(ch - rune('a'))
		if distance[n] != d {
			return false
		}
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
func isStrictlyPalindromic(n int) bool {
	for b := 2; b <= n-2; b++ {
		var s = tok(n, b)
		s = strings.TrimLeft(s, "0")
		if s != reverse(s) {
			return false
		}
	}
	return true
}

func tok(n int, k int) string {
	if n == 0 {
		return "0"
	}
	var s = n % k
	return tok(n/k, k) + string(s)
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
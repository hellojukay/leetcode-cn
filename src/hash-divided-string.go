func stringHash(s string, k int) string {
	var result string
	for i := 0; i < len(s); i = i + k {
		var c int
		for j := 0; j < k && i+j < len(s); j++ {
			c = c + int(byte(s[i+j])-byte('a'))
		}
		c = c % 26
		result = result + string(byte(c)+byte('a'))
	}
	return result
}
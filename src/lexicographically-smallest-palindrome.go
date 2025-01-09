func makeSmallestPalindrome(s string) string {
	var buffer = []byte(s)
	var i, j int
	for i, j = 0, len(buffer)-1; i < j; i, j = i+1, j-1 {
		if buffer[i] > buffer[j] {
			buffer[i] = buffer[j]
		} else {
			buffer[j] = buffer[i]
		}
	}
	return string(buffer)
}
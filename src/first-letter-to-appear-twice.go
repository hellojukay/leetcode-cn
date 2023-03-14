func repeatedCharacter(s string) byte {
	var hash = make(map[byte]bool)
	for _, ch := range []byte(s) {
		if _, ok := hash[ch]; ok {
			return ch
		}
		hash[ch] = true
	}
	return byte('a')
}
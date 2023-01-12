package main

func percentageLetter(s string, letter byte) int {
	var length = len(s)
	var count int
	for _, ch := range []byte(s) {
		if ch == letter {
			count++
		}
	}
	return count * 100 / length
}

func scoreOfString(s string) int {
	if len(s) == 2 {
		return abs(int(s[0]) - int(s[1]))
	}
	if len(s) <= 1 {
		return 0
	}
	return abs(int(s[0])-int(s[1])) + scoreOfString(s[1:])
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

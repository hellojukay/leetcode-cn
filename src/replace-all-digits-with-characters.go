func replaceDigits(s string) string {
	if len(s) == 0 {
		return ""
	}
	var result string
	i := 0
	j := 1
	for j < len(s) {
		var n = int(([]rune(s))[j] - rune('0'))
		result = result + fmt.Sprintf("%c", ([]rune(s))[i]) + shift(([]rune(s))[i], n)
		i = i + 2
		j = j + 2
	}
	if len(s)%2 == 1 {
		result = result + fmt.Sprintf("%c", []rune(s)[len(s)-1])
	}
	return result
}

func shift(c rune, n int) string {
	return fmt.Sprintf("%c", rune(c+rune(n)))
}
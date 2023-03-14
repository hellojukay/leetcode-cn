func removeVowels(s string) string {
	var result []rune
	for _, ch := range s {
		switch ch {
		case 'a':
			continue
		case 'o':
			continue
		case 'e':
			continue
		case 'i':
			continue
		case 'u':
			continue
		default:
			result = append(result, ch)

		}
	}
	return string(result)
}
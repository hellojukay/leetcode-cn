//https://leetcode-cn.com/problems/check-if-all-characters-have-equal-number-of-occurrences/
func areOccurrencesEqual(s string) bool {
	var hash = make(map[rune]int)
	for _, ch := range s {
		hash[ch]++
	}
	var c = ([]rune(s))[0]
	var count = hash[c]
	for _, v := range hash {
		if v != count {
			return false
		}
	}
	return true
}

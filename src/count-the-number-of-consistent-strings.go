//https://leetcode-cn.com/problems/count-the-number-of-consistent-strings/
func countConsistentStrings(allowed string, words []string) int {
	var hash = make(map[rune]bool)
	for _, ch := range allowed {
		hash[ch] = true
	}
	var count = 0
	var in = func(str string) bool {
		for _, ch := range str {
			if ok, _:= hash[ch];ok {
				continue
			}else {
				return false
			}
		}
		return true
	}
	for _, str := range words {
		if(in(str)) {
			count++
		}
	}
	return count
}
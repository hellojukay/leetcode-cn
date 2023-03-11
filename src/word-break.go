package main

func wordBreak(s string, wordDict []string) bool {
	var hash = make(map[string]bool)
	for _, word := range wordDict {
		hash[word] = true
	}
	return check([]rune(s), hash)
}

func check(s []rune, hash map[string]bool) bool {
	if _, ok := hash[string(s)]; ok {
		return true
	}
	for i := 1; i < len(s); i++ {
		sub := s[:i]
		if _, ok := hash[string(sub)]; ok {
			if check(s[i:], hash) {
				return true
			}
		}
	}
	return false
}

func main() {
	println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}

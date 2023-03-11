package main

func wordBreak(s string, wordDict []string) []string {
	var hash = make(map[string]bool)
	for _, word := range wordDict {
		hash[word] = true
	}
	var result = make([]string, 0)
	check(&result, "", []rune(s), hash)
	return result
}

func check(result *[]string, prefix string, s []rune, hash map[string]bool) {
	if len(s) == 0 {
		*result = append(*result, prefix)
		return
	}
	for i := 1; i < len(s); i++ {
		sub := s[:i]
		if _, ok := hash[string(sub)]; ok {
			if prefix == "" {
				check(result, string(sub), s[i:], hash)
			} else {
				check(result, prefix+" "+string(sub), s[i:], hash)
			}
		}
	}
	if _, ok := hash[string(s)]; ok {
		if prefix == "" {
			*result = append(*result, string(s))

		} else {
			*result = append(*result, prefix+" "+string(s))
		}
		return
	}
}

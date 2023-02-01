package main

func decodeMessage(key string, message string) string {
	var cache = make(map[rune]rune)
	var c = rune('a')
	for _, ch := range key {
		if c > rune('z') {
			break
		}
		if ch == rune(' ') {
			continue
		}
		if _, ok := cache[ch]; ok {
			continue
		}
		cache[ch] = c
		c = c + 1
	}
	var result []rune
	for _, ch := range message {
		v, ok := cache[ch]
		if ok {
			result = append(result, v)
		} else {
			result = append(result, ch)
		}
	}
	return string(result)
}

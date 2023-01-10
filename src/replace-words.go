package main

import "strings"

// https://leetcode.cn/problems/replace-words/
func replaceWords(dictionary []string, sentence string) string {
	var words = strings.Split(sentence, " ")
	var cache = make(map[string]bool)
	var max = 0
	for _, prefix := range dictionary {
		cache[prefix] = true
		var length = len(prefix)
		if length > max {
			max = len(prefix)
		}
	}
	for index, word := range words {
		var length = len(word)
		if length == 1 {
			continue
		}
		for i := 1; i < length; i++ {
			// 这里判断一下，如果前缀长度已经超过了字典里面最长的字符串了，那么比较就没有意义了
			if i > max {
				break
			}
			prefix := string([]rune(word[:i]))

			_, ok := cache[prefix]
			if ok {
				words[index] = prefix
				break
			}
		}
	}
	return strings.Join(words, " ")
}

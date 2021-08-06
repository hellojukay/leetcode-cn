package main

//https://leetcode-cn.com/problems/merge-strings-alternately/
func mergeAlternately(word1 string, word2 string) string {
	return term(word1,word2,"")
}


func term(str1 string, str2 string, result string) string {
	if str2 == "" {
		return result + str1
	}
	if str1 == "" {
		return result + str2
	}
	return term(string([]rune(str1)[1:]), string([]rune(str2)[1:]), result+string([]rune(str1)[0])+string([]rune(str2)[0]))
}

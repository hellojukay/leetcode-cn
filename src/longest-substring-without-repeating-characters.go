package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/hua-dong-chuang-kou-by-powcai/
func lengthOfLongestSubstring(s string) int {
	var length = len(s)
	if length == 0 || length == 1 {
		return length
	}
	var find_ch = func(src string, c rune) int {
		for index, ch := range src {
			if ch == c {
				return index
			}
		}
		return -1
	}
	var window string
	var max = 0
	for _, ch := range s {
		index := find_ch(window, ch)
		if index >= 0 {
			n := window[index+1:]
			println(window, index, n)
			window = n
		}
		window = window + string(ch)
		if len(window) > max {
			max = len(window)
		}
	}
	return max
}

func main() {
	var max = lengthOfLongestSubstring("abcabcbb")
	println(max)
}

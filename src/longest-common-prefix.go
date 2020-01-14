package main

func main() {
	var result = longestCommonPrefix([]string{"flower", "flow", "flight"})
	println(result)
}

// https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	var length = len(strs)
	if length == 0 {
		return ""
	}
	var result = strs[0]
	for i := 0; i < length; i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		result = xor(result, strs[i])
	}
	return result
}

func xor(src string, dest string) string {
	if src == "" {
		return ""
	}
	if dest == "" {
		return ""
	}
	if src == dest {
		return src
	}
	if len(src) > len(dest) {
		return xor(dest, src)
	}
	var min_length = len(src)
	for i := 0; i < min_length; i++ {
		if src[i] != dest[i] {
			if i == 0 {
				return ""
			}
			return src[:i]
		}
	}
	return src
}

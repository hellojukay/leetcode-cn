//https://leetcode-cn.com/problems/merge-strings-alternately/
func exchangeBits(n int) int {
	return term(n, 0, 0)
}

func sqre(n int, c int) int {
	if n == 0 {
		return 0
	}
	if c == 0 {
		return 1
	}
	return 1 << c
}
func term(n int, count int, result int) int {
	if count > 31 {
		return result
	}
	var low = n % 2
	var tmp = n >> 1
	var high = tmp % 2
	tmp = tmp >> 1
	return term(tmp, count+2, result+sqre(high, count)+sqre(low, count+1))
}


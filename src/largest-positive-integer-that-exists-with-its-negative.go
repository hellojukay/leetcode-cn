// https://leetcode.cn/problems/largest-positive-integer-that-exists-with-its-negative/
func findMaxK(nums []int) int {
	var result = -1
	var cache = make(map[int]bool)
	for _, n := range nums {
		if _, ok := cache[n*-1]; ok {
			if abs(n) > result {
				result = abs(n)
			}
		}
		cache[n] = true
	}
	return result
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
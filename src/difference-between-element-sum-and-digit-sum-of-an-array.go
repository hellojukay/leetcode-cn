func differenceOfSum(nums []int) int {
	s1 := sum(nums)
	var s2 int
	for _, n := range nums {
		s2 = s2 + sum(toArray(n))
	}
	return abs(s1 - s2)
}

func toArray(n int) []int {
	var result []int
	for n != 0 {
		c := n % 10
		result = append(result, c)
		n = n / 10
	}
	return result
}

func sum(arra []int) int {
	var result int
	for _, n := range arra {
		result = result + n
	}
	return result
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
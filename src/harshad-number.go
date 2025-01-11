func sumOfTheDigitsOfHarshadNumber(x int) int {
	var sum int
	var c int
	var n = x
	for n > 0 {
		c = n % 10
		sum += c
		n = n / 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}

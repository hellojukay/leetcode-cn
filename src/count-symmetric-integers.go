func countSymmetricIntegers(low int, high int) int {
	var count int
	for n := low; n <= high; n++ {
		if isSymmetric(n) {
			count++
		}
	}
	return count
}

func isSymmetric(n int) bool {
	var s = fmt.Sprintf("%d", n)
	if len(s)%2 != 0 {
		return false
	}
	var left = s[:len(s)/2]
	var right = s[len(s)/2:]
	return sum(left) == sum(right)
}

func sum(s string) int {
	var sum int
	for _, c := range s {
		sum += int(c - '0')
	}
	return sum
}

func count_zero(n int, count int) int {
	for {
		if n == 0 {
			break
		}
		if n%2 == 1 {
			count = count + 1
		}
		n = n >> 1
	}
	return count
}
func compare(a int, b int) bool {
	count_a := count_zero(a, 0)
	count_b := count_zero(b, 0)
	if count_a == count_b {
		return a < b
	}
	return count_a < count_b
}
func sortByBits(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var left []int
	var rigth []int
	var head = arr[0]
	for i := 1; i < len(arr); i++ {
		if compare(arr[i], head) {
			left = append(left, arr[i])
		} else {
			rigth = append(rigth, arr[i])
		}
	}
	result := sortByBits(left)
	result = append(result, head)
	return append(result, sortByBits(rigth)...)
}

func digitCount(num string) bool {
	var hash = make(map[int]int)
	for _, n := range []rune(num) {
		count, ok := hash[int(n-rune('0'))]
		if ok {
			count = count + 1
		} else {
			count = 1
		}
		hash[int(n-rune('0'))] = count
	}
	for i := 0; i < len(num); i++ {
		n, ok := hash[i]

		var tmp = int([]rune(num)[i] - rune('0'))
		if !ok && tmp != 0 {
			return false
		}
		if tmp != n {
			return false
		}
	}
	return true

}
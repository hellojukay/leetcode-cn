package main

func isHappy(n int) bool {
	var dict = make(map[int]bool)
	for {
		if n == 1 {
			break
		}
		dict[n] = true
		var tmp = 0
		for n != 0 {
			tmp = tmp + (n%10)*(n%10)
			n = n / 10
		}
		n = tmp
		if dict[tmp] {
			return false
		}
		dict[tmp] = true
		println(tmp)
	}
	return true
}

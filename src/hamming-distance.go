package main

func hammingDistance(x int, y int) int {
	var result = 0
	for !(x == 0 && y == 0) {
		if (x % 2) != (y % 2) {
			result++
		}
		x = x / 2
		y = y / 2
	}
	return result
}

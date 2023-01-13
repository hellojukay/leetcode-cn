package main

import "strconv"

func maximumValue(strs []string) int {
	var max = 0
	for _, str := range strs {
		var n = toInt(str)
		if n > max {
			max = n
		}
	}
	return max
}

func toInt(str string) int {
	number, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return len(str)
	}
	return int(number)
}

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var flag = false
	if x < 0 {
		x = -1 * x
		flag = true
	}
	var str = fmt.Sprintf("%d", x)
	var min = fmt.Sprintf("%d", math.MinInt32)
	var max = fmt.Sprintf("%d", math.MaxInt32)
	var tmp = ""
	for _, ch := range str {
		tmp = string(ch) + tmp
	}

	str = tmp
	tmp = strings.TrimLeft(tmp, "0")
	if flag {
		str = "-" + str
	}
	if !flag && len(str) > len(max) {
		return 0
	}
	if !flag && len(str) == len(max) && str > max {
		return 0
	}
	if flag && len(str) > len(min) {
		return 0
	}
	if flag && len(str) == len(min) && str > min {
		return 0
	}
	str = tmp
	if flag {
		str = "-" + tmp
	}
	str = strings.TrimLeft(str, "0")
	var y, _ = strconv.Atoi(str)
	return y
}

func main() {
	println(math.MaxInt32)
	println(math.MinInt32)
	println(reverse(10))
	println(reverse(-10))
	println(reverse(-1230))
	println(reverse(math.MaxInt32))
	println(reverse(math.MinInt32))
	println(reverse(-1463847412))
}

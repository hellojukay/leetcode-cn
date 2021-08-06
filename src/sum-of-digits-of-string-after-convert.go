
//https://leetcode-cn.com/problems/sum-of-digits-of-string-after-convert/
func char2int(n rune) int {
	tmp := int(n) - int('a') + 1
	println(tmp)
	return tmp
}

func string2int(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func term(number_string string, k int) int {
	if k == 0 {
		return string2int(number_string)
	}
	var sum = 0
	for _, ch := range number_string {
		sum = sum + string2int(string(ch))
	}
	return term(fmt.Sprintf("%d", sum), k-1)
}
func getLucky(s string, k int) int {
	var tmp string
	for _, ch := range s {
		tmp = tmp + fmt.Sprintf("%d", char2int(ch))
	}
	return term(tmp, k)
}

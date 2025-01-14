func removeTrailingZeros(num string) string {
	if len(num) == 0 {
		return num
	}
	if num[len(num)-1] == '0' {
		return removeTrailingZeros(num[:len(num)-1])
	}
	return num
}
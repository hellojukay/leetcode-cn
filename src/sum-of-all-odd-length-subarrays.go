func sumOddLengthSubarrays(arr []int) int {
	var s int
	for i := 0; i <= len(arr)-1; i++ {
		for length := 1; length <= len(arr); length += 2 {
			if i+length > len(arr) {
				break
			}
			sub := arr[i : i+length]
			s += sum(sub)
		}
	}
	return s
}

func sum(arr []int) int {
	var result int
	for _, n := range arr {
		result += n
	}
	return result
}
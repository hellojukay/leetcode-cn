func exchange(nums []int) []int {
	var result = make([]int, len(nums))
	var i = 0
	var j = len(nums) - 1
	for _, n := range nums {
		if n%2 == 0 {
			result[j] = n
			j--
		} else {
			result[i] = n
			i++
		}
	}
	return result
}
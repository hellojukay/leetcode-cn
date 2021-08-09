
func compare(a, b int, count map[int]int) bool {
	count_b := count[b]
	count_a := count[a]
	if count_a == count_b {
		return a > b
	}
	return count_a < count_b
}

func sort(nums []int, count map[int]int) []int {
	if len(nums) < 2 {
		return nums
	}
	var pat = nums[0]
	var left []int
	var right []int
	for i := 1; i < len(nums); i++ {
		if compare(nums[i], pat, count) {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	result := sort(left, count)
	result = append(result, pat)
	return append(result, sort(right, count)...)
}
func frequencySort(nums []int) []int {
	var count = make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	return sort(nums, count)
}

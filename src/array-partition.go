func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var sum int
	for index, n := range nums {
		if index%2 == 0 {
			sum += n
		}
	}
	return sum
}
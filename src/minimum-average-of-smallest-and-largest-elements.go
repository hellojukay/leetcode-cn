func minimumAverage(nums []int) float64 {
	//slice sort nums
	sort.Ints(nums)
	var farray []float64
	for len(nums) > 0 {
		//remove first and last element
		mid := float64(nums[0]+nums[len(nums)-1]) / 2
		farray = append(farray, mid)
		nums = nums[1 : len(nums)-1]
	}
	// sort farray
	sort.Float64s(farray)
	return farray[0]
}
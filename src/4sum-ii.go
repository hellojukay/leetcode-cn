func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var result int
	var hash1 = make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			n, ok := hash1[a+b]
			if ok {
				hash1[a+b] = n + 1
			} else {
				hash1[a+b] = 1
			}
		}
	}
	var hash2 = make(map[int]int)
	for _, a := range nums3 {
		for _, b := range nums4 {
			n, ok := hash2[a+b]
			if ok {
				hash2[a+b] = n + 1
			} else {
				hash2[a+b] = 1
			}
		}
	}
	for a, count1 := range hash1 {
		v, ok := hash2[-a]
		if ok {
			result = result + count1*v
		}
	}
	return result
}

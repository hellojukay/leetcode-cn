func heightChecker(heights []int) int {
	// sort int slices height in ascending order
	var tmp = mycopy(heights)
	sort.Ints(tmp)
	var count int
	for i := 0; i < len(heights); i++ {
		if heights[i] != tmp[i] {
			count++
		}
	}
	return count
}

func mycopy(src []int) []int {
	var s []int
	for _, n := range src {
		s = append(s, n)
	}
	return s
}
func peakIndexInMountainArray(arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[i+1] {
			continue
		}
		return i
	}
	return -1
}
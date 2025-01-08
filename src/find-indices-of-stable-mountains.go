func stableMountains(height []int, threshold int) []int {
	var result []int
	for i := 0; i < len(height); i++ {
		if i-1 >= 0 && height[i-1] > threshold {
			result = append(result, i)
		}
	}
	return result
}
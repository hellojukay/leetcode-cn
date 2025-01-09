func countTestedDevices(batteryPercentages []int) int {
	var result int
	for i := 0; i < len(batteryPercentages); i++ {
		if batteryPercentages[i]-result > 0 {
			result++
		}
	}
	return result
}
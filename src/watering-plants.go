func wateringPlants(plants []int, capacity int) int {
	var res = capacity
	var count = 0
	for i, p := range plants {
		if i == 0 {
			count = 1
			res = capacity - p
			continue
		}
		if res >= p {
			res = res - p
			count = count + 1
			continue
		}
		count = count + 1 + i*2
		res = capacity - p
	}
	return count
}
package main

func findLonely(nums []int) []int {
	var dict = make(map[int]int)
	for _, n := range nums {
		v, ok := dict[n]
		if ok {
			v++
		} else {
			v = 1
		}
		dict[n] = v
	}
	var result []int
	for _, n := range nums {
		count := dict[n]
		if count > 1 {
			continue
		}
		if _, ok := dict[n-1]; !ok {
			if _, ok = dict[n+1]; !ok {
				result = append(result, n)
			}
		}
	}
	return result
}

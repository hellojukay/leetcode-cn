package main

func duplicateZeros(arr []int) {
	var dest = make([]int, len(arr))
	copy(dest, arr)
	index := 0
	for _, n := range dest {
		if index == len(dest) {
			break
		}
		if n == 0 {
			arr[index] = 0
			index = index + 1
			if index == len(dest) {
				break
			} else {
				arr[index] = 0
			}
		} else {
			arr[index] = n
		}
		index = index + 1
	}
}

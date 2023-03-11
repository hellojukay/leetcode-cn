package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var result = make([][]int, 0)
	sort.Ints(candidates)
	combine(candidates, []int{}, &result, target)
	return result
}

func combine(candidates []int, prefix []int, result *[][]int, target int) {
	if target == 0 && len(prefix) != 0 {
		*result = append(*result, prefix)
		return
	}
	if candidates[0] > target {
		return
	}
	for _, n := range candidates {
		if n <= target {
			combine(candidates, append(getCopy(prefix), n), result, target-n)
		}
	}
}

func getCopy(src []int) []int {
	var buffer = make([]int, len(src))
	copy(buffer, src)
	return buffer
}

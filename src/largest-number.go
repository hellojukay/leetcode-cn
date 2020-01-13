package main

import "fmt"

//https://leetcode-cn.com/problems/largest-number/
func find_max(nums []int) string {
	var length = len(nums)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return fmt.Sprintf("%d", nums[0])
	}
	var max = ""
	var pos = 0
	for index, n := range nums {
		str := fmt.Sprintf("%d", n)
		if max+str <= str+max {
			max = str
			pos = index
		}
	}
	var arr []int
	for index, n := range nums {
		if index != pos {
			arr = append(arr, n)
		}
	}
	return max + find_max(arr)
}
func largestNumber(nums []int) string {
	var max = find_max(nums)
	var flag = false
	for index := range max {
		if max[index] != '0' {
			flag = true
		}
	}
	if !flag {
		return "0"
	}
	return max
}

func main() {
	println(largestNumber([]int{3, 30, 34, 5, 9}))
	println(largestNumber([]int{0, 0}))
	println(largestNumber([]int{1, 0, 0}))
}

package main

import "fmt"

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge
	var size = len(nums1) + len(nums2)
	var arr = make([]int, size)
	// num1 index
	var i = 0
	// num2 index
	var j = 0
	var index = 0
	for {
		if i >= len(nums1) && j >= len(nums2) {
			break
		}
		if i >= len(nums1) {
			arr[index] = nums2[j]
			j++
			index++
			continue
		}
		if j >= len(nums2) {
			arr[index] = nums1[i]
			index++
			i++
			continue
		}
		if nums1[i] < nums2[j] {
			arr[index] = nums1[i]
			index++
			i++
			continue
		} else {
			arr[index] = nums2[j]
			index++
			j++
		}
	}
	fmt.Printf("%v", arr)
	// get mdiean
	if size%2 == 1 {
		return float64(arr[((size+1)/2)-1])
	}
	return float64(arr[(size/2)-1])/2.0 + float64(arr[size/2])/2.0
}

func main() {
	findMedianSortedArrays([]int{1, 3}, []int{2})
}

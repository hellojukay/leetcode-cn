package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//  https://leetcode-cn.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	var tmp = make([]int, m)
	for i := 0; i < m; i++ {
		tmp[i] = nums1[i]
	}
	var i, j = 0, 0
	var index = 0
	for {
		if i >= m && j >= n {
			break
		}
		if i >= m {
			nums1[index] = nums2[j]
			j++
			index++
			continue
		}
		if j >= n {
			nums1[index] = tmp[i]
			i++
			index++
			continue
		}
		if tmp[i] < nums2[j] {
			nums1[index] = tmp[i]
			i++
			index++
			continue
		} else {
			nums1[index] = nums2[j]
			j++
			index++
		}
	}
}

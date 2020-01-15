package main

import "fmt"

func main() {
	var arr = []int{4, 5}
	fmt.Printf("%v\n", arr)
	var n = removeElement(arr, 5)
	fmt.Printf("%v\n", arr)
	println(n)
}

// https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	var count = 0
	var length = len(nums)
	if length == 0 {
		return 0
	}
	var i = 0
	var j = length - 1
	for {
		if nums[i] == val {
			count++
			for {
				if j <= i {
					break
				}
				if val != nums[j] {
					nums[i] = nums[j]
					j--
					break
				}
				count++
				j--
			}
		}
		i++
		if j < i {
			break
		}
	}
	return length - count
}

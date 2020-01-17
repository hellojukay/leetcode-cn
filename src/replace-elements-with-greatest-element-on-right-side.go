package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/
func replaceElements(arr []int) []int {
	var length = len(arr)
	var max = -1
	var max_index = -1
	for i := 0; i < length; i++ {
		if i < max_index {
			arr[i] = max
			continue
		}
		if i+1 < length {
			max = arr[i+1]
		} else {
			arr[i] = -1
			break
		}
		for index := i + 1; index < length; index++ {
			if max < arr[index] {
				max = arr[index]
				max_index = index
			}
		}
		arr[i] = max
	}
	return arr
}

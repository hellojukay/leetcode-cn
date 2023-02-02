package main

func sortPeople(names []string, heights []int) []string {
	var slices []*HeightIndex
	for index := range heights {
		slices = append(slices, &HeightIndex{Height: heights[index], Name: names[index]})
	}
	slices = sort(slices)
	var result []string
	for _, heiHeightIndex := range slices {
		result = append(result, heiHeightIndex.Name)
	}
	return result
}

type HeightIndex struct {
	Height int
	Name   string
}

func sort(heights []*HeightIndex) []*HeightIndex {
	if len(heights) <= 1 {
		return heights
	}
	var left []*HeightIndex
	var right []*HeightIndex
	var mid = heights[0]
	for i := 1; i < len(heights); i++ {
		if heights[i].Height > mid.Height {
			left = append(left, heights[i])
		} else {
			right = append(right, heights[i])
		}
	}
	return append(append(sort(left), mid), sort(right)...)
}

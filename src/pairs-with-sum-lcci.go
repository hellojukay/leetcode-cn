//https://leetcode-cn.com/problems/pairs-with-sum-lcci/
func pairSums(nums []int, target int) [][]int {
	var m = make(map[int]int)
	for _, n := range nums{
		m[n] = m[n] + 1
	}
	var result [][]int
	for _, n := range nums{
		if m[n] == 0 || m[target-n] == 0{
			continue
		}
		if m[target-n] > 0 {
			if n == target-n {
				if m[n] > 1 {
					result = append(result,[]int{n,n})
					m[n] = m[n] -2
				} 
			}else {
				result = append(result,[]int{n,target-n})
				m[target-n] = m[target-n]-1
				m[n] = m[n]-1
			}
		}
	}
	return result
}
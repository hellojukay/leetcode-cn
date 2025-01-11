func generateKey(num1 int, num2 int, num3 int) int {
	var digits1 []int = intToDigits(num1)
	var digits2 []int = intToDigits(num2)
	var digits3 []int = intToDigits(num3)
	var result int
	result = getMin([]int{digits1[0], digits2[0], digits3[0]}) * 1000
	result += getMin([]int{digits1[1], digits2[1], digits3[1]}) * 100
	result += getMin([]int{digits1[2], digits2[2], digits3[2]}) * 10
	result += getMin([]int{digits1[3], digits2[3], digits3[3]})
	return result
}

func intToDigits(num int) []int {
	var digits []int = []int{0, 0, 0, 0}
	digits[0] = num / 1000
	digits[1] = (num % 1000) / 100
	digits[2] = (num % 100) / 10
	digits[3] = num % 10
	return digits
}

func getMin(nums []int) int {
	var min int = nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}
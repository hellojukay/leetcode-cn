func largestGoodInteger(num string) string {
	var array = []string{"000", "111", "222", "333", "444", "555", "666", "777", "888", "999"}
	for i := 9; i >= 0; i-- {
		if strings.Contains(num, array[i]) {
			return array[i]
		}
	}
	return ""
}
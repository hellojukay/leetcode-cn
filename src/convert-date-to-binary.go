func convertDateToBinary(date string) string {
	array := strings.Split(date, "-")
	year, _ := strconv.Atoi(array[0])
	month, _ := strconv.Atoi(array[1])
	day, _ := strconv.Atoi(array[2])
	return intToBinary(year) + "-" + intToBinary(month) + "-" + intToBinary(day)
}

func intToBinary(num int) string {
	return strconv.FormatInt(int64(num), 2)
}
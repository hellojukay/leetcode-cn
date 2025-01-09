func titleToNumber(columnTitle string) int {
	if len(columnTitle) == 1 {
		return int(columnTitle[0] - 'A' + 1)
	}
	return int(columnTitle[len(columnTitle)-1]-'A'+1) + 26*titleToNumber(columnTitle[:len(columnTitle)-1])
}

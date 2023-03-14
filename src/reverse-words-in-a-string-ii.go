func reverseWords(s []byte) {
	if len(s) <= 1 {
		return
	}
	reverse(s)

	var flag bool
	var start int
	var end int
	for end = start + 1; end < len(s); end++ {
		if !flag {
			if s[end] == byte(' ') {
				flag = true
				reverse(s[start:end])
			}
		} else {
			if s[end] != byte(' ') {
				start = end
				flag = false
			}
		}
	}

	reverse(s[start:end])

}

func reverse(s []byte) {
	if len(s) <= 1 {
		return
	}
	for i := 0; i < len(s)-1-i; i++ {
		tmp := s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = tmp
	}
}
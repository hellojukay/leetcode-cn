func rearrangeCharacters(s string, target string) int {
	if len(s) < len(target) {
		return 0
	}
	var max = 0
	var hash1 = make(map[rune]int)
	for _, ch := range target {
		v, ok := hash1[ch]
		if ok {
			v++
		} else {
			v = 1
		}
		hash1[ch] = v
	}
	var hash = make(map[rune]int)
	for _, ch1 := range target {
		var count = 0
		for _, ch := range s {
			if ch1 == ch {
				count++
			}
		}
		v, ok := hash1[ch1]
		if !ok {
			count = 0
		} else {
			count = count / v
		}
		hash[ch1] = count
		if count > max {
			max = count
		}
	}
	var min = max
	for _, count := range hash {
		if count < min {
			min = count
		}
	}
	return min
}

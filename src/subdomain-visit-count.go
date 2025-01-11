func subdomainVisits(cpdomains []string) []string {
	var hash = make(map[string]int)
	for _, domain := range cpdomains {
		countS := strings.Split(domain, " ")[0]
		//convert countS to int
		count, _ := strconv.Atoi(countS)
		domainS := strings.Split(domain, " ")[1]
		array := strings.Split(domainS, ".")
		for i := 0; i < len(array); i++ {
			domain := strings.Join(array[i:], ".")
			hash[domain] += count
		}
	}
	var result []string
	for key, value := range hash {
		result = append(result, fmt.Sprintf("%d %s", value, key))
	}
	return result
}
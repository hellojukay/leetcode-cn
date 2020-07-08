package main

func findTheDifference(s string, t string) byte {
	var result byte
	for _, ch := range s {
		result = result ^ byte(ch)
	}
	for _, ch := range t {
		result = result ^ byte(ch)
	}
	return result
}

package main

func main() {
	println(addBinary("11", "1"))
	println(addBinary("1010", "1011"))
}

// https://leetcode-cn.com/problems/add-binary/
func addBinary(a string, b string) string {
	if len(b) > len(a) {
		return addBinary(b, a)
	}
	var c = "0"
	var result = ""
	for i := 0; i < len(a); i++ {
		if i < len(b) {
			ch, tmp := sum(string(a[len(a)-i-1]), string(b[len(b)-i-1]), c)
			result = ch + result
			c = tmp
		} else {
			ch, tmp := sum(string(a[len(a)-i-1]), "0", c)
			result = ch + result
			c = tmp
		}
	}
	if c == "1" {
		result = c + result
	}
	return result
}
func sum(a string, b string, c string) (string, string) {
	if c == "1" {
		if a != b {
			return "0", "1"
		}
		if a == b && a == "1" {
			return "1", "1"
		}
		return "1", "0"
	}
	if a != b {
		return "1", "0"
	}
	if a == b && a == "1" {
		return "0", "1"
	}
	return "0", "0"
}

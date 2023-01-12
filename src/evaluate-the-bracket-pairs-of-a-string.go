package main

import "bytes"

func evaluate(s string, knowledge [][]string) string {
	var hash = make(map[string]string)
	for _, arr := range knowledge {
		hash[arr[0]] = arr[1]
	}
	var in = false
	var key string
	var buffer bytes.Buffer
	for _, ch := range []rune(s) {
		if ch == rune('(') {
			in = true
			continue
		}
		if in && ch != rune(')') {
			key = key + string([]rune{ch})
			continue
		}
		if ch == rune(')') {
			v, ok := hash[key]
			if ok {
				buffer.Write([]byte(v))
			} else {
				buffer.WriteRune(rune('?'))
			}
			key = ""
			in = false
			continue
		}
		buffer.WriteRune((ch))
	}
	return buffer.String()
}

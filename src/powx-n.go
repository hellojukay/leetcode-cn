package main

func myPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		return 1 / myPow(x, -1*n)
	}
	if n%2 == 0 {
		var tmp = myPow(x, n/2)
		return tmp * tmp
	}
	var tmp = myPow(x, (n-1)/2)
	return tmp * tmp * x
}
func main() {
	println(myPow(2.0, 10))
}

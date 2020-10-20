package main

import "fmt"

func myPow(x float64, n int) float64 {
	if x == 0 {
		return x
	}
	if n == 0 {
		return 1
	}
	result := 1.0
	if n < 0 {
		x = 1.0 / x
		n = -n
	}
	for n != 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}
	return result
}

func main() {
	fmt.Print(myPow(34.00515, -3))
}

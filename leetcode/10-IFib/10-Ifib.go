package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	intArray := make([]int, n+1)
	intArray[0] = 0
	intArray[1] = 1
	if n > 1 {
		for i := 1; i < n; i++ {
			intArray[i+1] = (intArray[i] + intArray[i-1]) % 1000000007
		}
	}
	return intArray[n]
}

func main() {
	fmt.Print(fib(45))
}

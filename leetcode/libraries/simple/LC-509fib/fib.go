package main

var chache = make(map[int]int)

func fib(n int) int {
	if n == 0 {
		chache[0] = 0
		return 0
	}
	if n == 1 {
		chache[1] = 1
		return 1
	}
	fir := 0
	sec := 0
	if value, ok := chache[n-1]; ok {
		fir = value
	} else {
		fir = fib(n - 1)
		chache[n-1] = fir
	}
	if value, ok := chache[n-2]; ok {
		sec = value
	} else {
		sec = fib(n - 2)
		chache[n-2] = sec
	}
	result := 0
	if value, ok := chache[n]; ok {
		return value
	} else {
		result = fir + sec
		chache[n] = result
		return result
	}
}

func main() {
	fib(2)
}

package main

func prefixesDivBy5(A []int) []bool {
	result := make([]bool, len(A))
	prev := 0
	for i, a := range A {
		prev = ((prev << 1) + a) % 5
		result[i] = (prev == 0)
	}
	return result
}

func main() {
	prefixesDivBy5([]int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0})
}

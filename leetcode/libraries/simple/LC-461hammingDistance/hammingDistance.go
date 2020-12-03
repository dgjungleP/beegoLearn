package main

import "fmt"

func hammingDistance(x int, y int) int {
	count := 0
	for x != 0 || y != 0 {
		if x%2 != y%2 {
			count++
		}
		x = x / 2
		y = y / 2
	}
	return count

}

func main() {
	fmt.Printf("%d", hammingDistance(1, 4))
}

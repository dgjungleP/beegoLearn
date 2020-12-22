package main

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

func main() {

}

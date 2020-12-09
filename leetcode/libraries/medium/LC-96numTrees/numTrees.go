package main

func numTrees(n int) int {
	catalan := 1
	for i := 0; i < n; i++ {
		catalan = catalan * (2 * (2*i + 1)) / (i + 1)
	}
	return catalan
}

func main() {

}

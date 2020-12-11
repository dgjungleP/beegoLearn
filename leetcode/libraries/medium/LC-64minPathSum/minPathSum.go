package main

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 {
				if j != 0 {
					grid[i][j] += grid[i][j-1]
				}
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
func min(numb1, numb2 int) int {
	if numb1 > numb2 {
		return numb2
	}
	return numb1
}
func main() {

}

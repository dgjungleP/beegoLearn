package main

import "fmt"

func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			} else if i > 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			} else if i > 0 && j > 0 {
				if grid[i][j-1] > grid[i-1][j] {
					grid[i][j] += grid[i][j-1]
				} else {
					grid[i][j] += grid[i-1][j]
				}
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	fmt.Print(maxValue([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}}))
}

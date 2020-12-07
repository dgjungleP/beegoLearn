package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	res := []int{}
	result := make([][]int, 0)
	var dfs func(int, int)
	dfs = func(index int, sum int) {
		if index >= len(candidates) {
			return
		}
		if sum == target {
			result = append(result, append([]int{}, res...))
			return
		} else if sum > target {
			return
		}
		if sum+candidates[index] <= target {
			res = append(res, candidates[index])
			dfs(index, sum+candidates[index])
			res = res[:len(res)-1]
		}
		dfs(index+1, sum)
	}
	sort(candidates)
	dfs(0, 0)
	return result
}

func sort(candidates []int) {
	for i := 0; i < len(candidates); i++ {
		for j := i; j < len(candidates); j++ {
			if candidates[j] < candidates[i] {
				candidates[j], candidates[i] = candidates[i], candidates[j]
			}
		}
	}
}

func main() {
	fmt.Print(combinationSum([]int{2, 7, 6, 3, 5, 1}, 9))
}

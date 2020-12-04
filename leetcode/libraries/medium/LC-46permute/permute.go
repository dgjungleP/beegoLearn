package main

import "fmt"

func permute(nums []int) [][]int {
	result := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for _, num := range nums {
			if visited[num] {
				continue
			}
			path = append(path, num)
			visited[num] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[num] = false
		}
	}
	dfs([]int{})
	return result
}

func main() {
	fmt.Print(permute([]int{1, 2, 3}))

}

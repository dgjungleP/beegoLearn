package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] ||
			(people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	for i, v := range people {
		copy(people[v[1]+1:i+1], people[v[1]:i+1])
		people[v[1]] = v
	}
	return people
}

func main() {
	reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}})
}

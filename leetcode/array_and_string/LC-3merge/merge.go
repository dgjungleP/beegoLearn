package main

import "fmt"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	intervals = sort(intervals)
	result := make([][]int, 0)
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if result[len(result)-1][1] >= intervals[i][0] {
			if result[len(result)-1][1] < intervals[i][1] {
				result[len(result)-1][1] = intervals[i][1]
			}
			if result[len(result)-1][0] >= intervals[i][0] {
				result[len(result)-1][0] = intervals[i][0]
			}
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}
func sort(intervals [][]int) [][]int {
	for i := 0; i < len(intervals); i++ {
		for j := i; j < len(intervals); j++ {
			if intervals[j][0] <= intervals[i][0] {
				temp := intervals[i]
				intervals[i] = intervals[j]
				intervals[j] = temp
			}
		}
	}
	return intervals
}

func main() {
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}

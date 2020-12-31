package main

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return 0
	}
	for i := 0; i < len(intervals); i++ {
		for j := i; j < len(intervals); j++ {
			if intervals[i][1] > intervals[j][1] {
				intervals[j], intervals[i] = intervals[i], intervals[j]
			}
		}
	}
	count := 0
	tail := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < tail {
			count++
		} else {
			tail = intervals[i][1]
		}
	}
	return count
}

func main() {
	eraseOverlapIntervals([][]int{[]int{1, 2}, []int{3, 4}, []int{2, 3}, []int{1, 3}})
}

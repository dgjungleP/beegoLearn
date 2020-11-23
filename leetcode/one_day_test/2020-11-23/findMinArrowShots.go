package main

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	if len(points) == 1 {
		return 1
	}
	for i := 0; i < len(points); i++ {
		for j := i; j < len(points); j++ {
			if points[i][0] >= points[j][0] {
				points[i], points[j] = points[j], points[i]
			}
		}
	}
	result := [][]int{points[0]}
	cover := false
	for i := 1; i < len(points); i++ {
		if points[i][0] > result[len(result)-1][1] || cover {
			result = append(result, points[i])
			cover = false
		} else if points[i][1] >= result[len(result)-1][1] {
			result[len(result)-1][1] = points[i][1]
			cover = true
		} else {
			cover = true
		}
	}
	return len(result)
}

/** error**/
func main() {
	findMinArrowShots([][]int{{1, 2}, {4, 5}, {1, 5}})
}

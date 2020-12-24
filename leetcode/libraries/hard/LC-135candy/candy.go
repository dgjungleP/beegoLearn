package main

func candy(ratings []int) int {
	left := make([]int, len(ratings))
	right := make([]int, len(ratings))
	left[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] <= ratings[i-1] {
			left[i] = 1
		} else {
			left[i] = left[i-1] + 1
		}
	}
	right[len(ratings)-1] = 1
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] <= ratings[i+1] {
			right[i] = 1
		} else {
			right[i] = right[i+1] + 1
		}
	}
	result := 0
	for i := 0; i < len(ratings); i++ {
		result += max(left[i], right[i])
	}
	return result
}
func max(left, right int) int {
	if left < right {
		return right
	}
	return left
}
func main() {
	candy([]int{1, 0, 2})
}

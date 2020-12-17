package main

func dailyTemperatures(T []int) []int {
	length := len(T)
	result := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temp := T[i]
		for len(stack) > 0 && temp > T[stack[len(stack)-1]] {
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[pre] = i - pre
		}
		stack = append(stack, i)
	}
	return result
}
func main() {
	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
}

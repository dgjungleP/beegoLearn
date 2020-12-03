package main

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		index := nums[i]
		if index < 0 {
			index = -index
		}
		index--
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	for index, value := range nums {
		if value > 0 {
			result = append(result, index+1)
		}
	}
	return result
}

func main() {
	findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
}

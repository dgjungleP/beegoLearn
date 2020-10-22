package main

func containsDuplicate(nums []int) bool {
	valueMap := make(map[int]int)

	for _, v := range nums {
		if _, ok := valueMap[v]; ok {
			return true
		} else {
			valueMap[v] = 1
		}
	}
	return false
}
func main() {

}

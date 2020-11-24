package main

func majorityElement(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums[len(nums)/2]
}

func main() {

}

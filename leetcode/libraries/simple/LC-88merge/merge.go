package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	up := m - 1
	down := n - 1
	index := m + n - 1
	for up >= 0 && down >= 0 {
		if nums1[up] > nums2[down] {
			nums1[index] = nums1[up]
			up--
		} else {
			nums1[index] = nums2[down]
			down--
		}
		index--
	}
	for down >= 0 {
		nums1[down] = nums2[down]
		down--
	}
}
func main() {
	merge([]int{0}, 0, []int{1}, 1)
}

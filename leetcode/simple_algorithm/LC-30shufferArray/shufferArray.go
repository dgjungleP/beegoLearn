package main

type Solution struct {
	baseArr []int
	showArr []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums,
		nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	this.showArr = this.baseArr
	return this.showArr
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {

	return this.showArr
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {

}

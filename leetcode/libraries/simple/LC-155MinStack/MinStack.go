package main

import "fmt"

type MinStack struct {
	data     []int
	min_data []int
	size     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:     make([]int, 0),
		min_data: make([]int, 0),
		size:     0}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if this.size != 0 {
		min := this.min_data[this.size-1]
		if x > min {
			x = min
		}
	}
	this.min_data = append(this.min_data, x)
	this.size++
}

func (this *MinStack) Pop() {
	this.size--
	this.data = this.data[:this.size]
	this.min_data = this.min_data[:this.size]
}

func (this *MinStack) Top() int {
	return this.data[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.min_data[this.size-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	x := 1
	obj := Constructor()
	obj.Push(x)
	obj.Pop()
	param3 := obj.Top()
	param4 := obj.GetMin()
	fmt.Printf("%d", param3)
	fmt.Printf("%d", param4)
}

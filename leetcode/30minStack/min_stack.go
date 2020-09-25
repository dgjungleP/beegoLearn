package main

import (
	"container/list"
	"fmt"
)

type MinStack struct {
	realStack      *list.List
	auxiliaryStack *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		realStack:      new(list.List),
		auxiliaryStack: new(list.List)}
}

func (this *MinStack) Push(x int) {
	if this.auxiliaryStack.Len() == 0 {
		this.auxiliaryStack.PushBack(x)
	}
	if this.realStack.Back() != nil {
		top := this.auxiliaryStack.Back().Value.(int)
		if top >= x {
			this.auxiliaryStack.PushBack(x)
		}
	}
	this.realStack.PushBack(x)
}

func (this *MinStack) Pop() {
	top := this.realStack.Back().Value.(int)
	min := this.auxiliaryStack.Back().Value.(int)
	if top == min {
		this.auxiliaryStack.Remove(this.auxiliaryStack.Back())
	}
	this.realStack.Remove(this.realStack.Back())

}

func (this *MinStack) Top() int {
	return this.realStack.Back().Value.(int)
}

func (this *MinStack) Min() int {
	return this.auxiliaryStack.Back().Value.(int)
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-1)
	fmt.Printf("%d\n", obj.Min())
	fmt.Printf("%d\n", obj.Top())
	obj.Pop()
	fmt.Printf("%d", obj.Min())
}

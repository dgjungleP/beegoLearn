package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	preStack  *list.List
	tailStack *list.List
}

func Constructor() CQueue {
	return CQueue{
		preStack:  list.New(),
		tailStack: list.New()}
}

func (this *CQueue) AppendTail(value int) {
	this.preStack.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.tailStack.Len() == 0 {
		for this.preStack.Len() > 0 {
			this.tailStack.PushBack(this.preStack.Remove(this.preStack.Back()))
		}
	}
	if this.tailStack.Len() != 0 {
		value := this.tailStack.Back()
		this.tailStack.Remove(value)
		return value.Value.(int)
	}
	return -1
}

func main() {
	var result int
	cqueue := Constructor()
	cqueue.AppendTail(1)
	cqueue.AppendTail(2)
	cqueue.AppendTail(3)
	result = cqueue.DeleteHead()
	fmt.Printf("%d\n", result)
	result = cqueue.DeleteHead()
	fmt.Printf("%d\n", result)
	result = cqueue.DeleteHead()
	fmt.Printf("%d\n", result)
	result = cqueue.DeleteHead()
	fmt.Printf("%d", result)
}

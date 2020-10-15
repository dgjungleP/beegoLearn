package main

import "fmt"

type MedianFinder struct {
	data []int
	size int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		data: make([]int, 0),
		size: 0}
}

func (this *MedianFinder) AddNum(num int) {
	this.data = append(this.data, num)
	this.size++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.size&1 == 1 {
		return float64(this.data[this.size/2])
	} else {
		return (float64(this.data[this.size/2]) + float64(this.data[this.size/2-1])) / 2
	}
}

func main() {
	test := Constructor()
	test.AddNum(6)
	fmt.Println(test.FindMedian())
	test.AddNum(10)
	fmt.Println(test.FindMedian())
	test.AddNum(2)
	fmt.Println(test.FindMedian())
	test.AddNum(6)
	fmt.Println(test.FindMedian())
	test.AddNum(5)
	fmt.Println(test.FindMedian())
	test.AddNum(0)
	fmt.Println(test.FindMedian())
	test.AddNum(6)
	fmt.Println(test.FindMedian())
	test.AddNum(3)
	fmt.Println(test.FindMedian())
	test.AddNum(1)
	fmt.Println(test.FindMedian())
	test.AddNum(0)
	fmt.Println(test.FindMedian())
	test.AddNum(0)
	fmt.Println(test.FindMedian())

}

package main

import "fmt"

type MinStack struct {
	top   int
	stack []int
}

func Constructor() MinStack {
	stack := []int{}
	s := MinStack{top: -1, stack: stack}
	return s
}

func (this *MinStack) Push(val int) {
	if this.top == -1 {
		this.top = 0
	} else {
		this.top++
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if this.top == -1 {
		return
	}
	this.stack = this.stack[:this.top]
	this.top--
}

func (this *MinStack) Top() int {
	return this.stack[this.top]
}

func (this *MinStack) GetMin() int {
	ret := this.stack[0]
	for i := 0; i <= this.top; i++ {
		if ret > this.stack[i] {
			ret = this.stack[i]
		}
	}
	return ret
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())

	obj.Pop()
	fmt.Println(obj.Top())

	fmt.Println(obj)
}

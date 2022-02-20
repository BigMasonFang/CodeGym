package main

import (
	"bytes"
)

type Stack struct {
	top   int
	stack []byte
}

func stackConstructor() Stack {
	stack := []byte{}
	s := Stack{top: -1, stack: stack}
	return s
}

func (this *Stack) Push(val byte) {
	if this.top == -1 {
		this.top = 0
	} else {
		this.top++
	}
	this.stack = append(this.stack, val)
}

func (this *Stack) Pop() {
	if this.top == -1 {
		return
	}
	this.stack = this.stack[:this.top]
	this.top--
}

func backspaceCompare(s string, t string) bool {
	stackS, stackT := stackConstructor(), stackConstructor()
	for _, v := range s {
		if string(v) == "#" {
			stackS.Pop()
		} else {
			stackS.Push(byte(v))
		}
	}
	for _, v := range t {
		if string(v) == "#" {
			stackT.Pop()
		} else {
			stackT.Push(byte(v))
		}
	}
	return bytes.Compare(stackS.stack, stackT.stack)
}

func main() {

}

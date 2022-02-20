package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	top   int
	stack []int
}

func stackConstructor() Stack {
	stack := []int{}
	s := Stack{top: -1, stack: stack}
	return s
}

func (this *Stack) Push(val int) {
	if this.top == -1 {
		this.top = 0
	} else {
		this.top++
	}
	this.stack = append(this.stack, val)
}

func (this *Stack) Pop() int {
	if this.top == -1 {
		return 0
	}
	ret := this.stack[this.top]
	this.stack = this.stack[:this.top]
	this.top--
	return ret
}

func (this *Stack) Top() int {
	return this.stack[this.top]
}

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		val, _ := strconv.Atoi(tokens[0])
		return val
	}
	stack := stackConstructor()

	for _, v := range tokens {
		switch string(v) {
		case "+":
			stack.Push(stack.Pop() + stack.Pop())
		case "-":
			m, n := stack.Pop(), stack.Pop()
			stack.Push(n - m)
		case "*":
			stack.Push(stack.Pop() * stack.Pop())
		case "/":
			m, n := stack.Pop(), stack.Pop()
			stack.Push(n / m)
		default:
			val, _ := strconv.Atoi(v)
			stack.Push(val)
		}
		fmt.Println(stack, stack.Top())
	}
	return stack.Pop()
}

func main() {
	test := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(test))
}

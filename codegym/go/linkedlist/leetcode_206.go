package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// iteratively
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// stack
	slice := make([]int, 0)
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	i := len(slice) - 1
	newHead := &ListNode{Val: slice[i]}
	newHead.Next = &ListNode{Val: slice[i-1]}
	rst := newHead
	for i > 0 {
		i--
		newHead.Next = &ListNode{Val: slice[i-1]}
		newHead = newHead.Next
	}
	return rst
}

// standard 20220305 redo
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev, curr := &ListNode{}, head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	var test *ListNode
	fmt.Println(test)
	test2 := &ListNode{}
	fmt.Println(test2)
}

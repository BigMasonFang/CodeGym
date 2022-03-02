package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	if p == nil {
		return head
	}
	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

func main() {
	n5 := &ListNode{Val: 3, Next: nil}
	n4 := &ListNode{Val: 3, Next: n5}
	n3 := &ListNode{Val: 2, Next: n4}
	n2 := &ListNode{Val: 1, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}

	rst := deleteDuplicates(n1)
	fmt.Println(rst)
}

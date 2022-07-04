package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := &ListNode{Val: 0, Next: head}
	p_prev := ret
	p2 := head.Next
	// head as p1
	for p2.Next != nil {
		if head.Val == p2.Val {
			for head.Val == p2.Val {
				p2 = p2.Next
				if p2 == nil {
					break
				}
			}
			head = p_prev
			head.Next = p2
			if p2 == nil {
				break
			}
		} else {
			p_prev = head
			head = head.Next
			p2 = p2.Next
		}
	}
	if p2 == nil {
		return ret.Next
	}
	if head.Val == p2.Val {
		head = p_prev
		head.Next = nil
	}
	return ret.Next
}

func main() {
	n4 := ListNode{Val: 2, Next: nil}
	n3 := ListNode{Val: 2, Next: &n4}
	n2 := ListNode{Val: 2, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}
	retH := deleteDuplicates(&n1)

	for retH.Next != nil {
		fmt.Println(retH)
		retH = retH.Next
	}
	fmt.Println(retH)
}

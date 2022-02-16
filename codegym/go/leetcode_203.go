package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	// find head
	for head.Val == val {
		head = head.Next
		if head == nil {
			return head
		}
	}

	p, previous := head, head

	for p.Next != nil {
		// this come first
		if p.Val != val {
			previous = p
		} else {
			previous.Next = p.Next
		}
		p = p.Next
	}

	if p.Val == val {
		previous.Next = nil
	}
	return head
}

func removeElements(head *ListNode, val int) *ListNode {
	// create a head
	new_head := &ListNode{Val: 0, Next: head}
	prev := new_head
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return new_head.Next
}

func main() {
	n5 := ListNode{Val: 2, Next: nil}
	n4 := ListNode{Val: 5, Next: &n5}
	n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}

	retH := *removeElements(&n1, 2)
	// retN := *removeElements(nil, 0)
	// fmt.Println(retH)

	for retH.Next != nil {
		fmt.Println(retH)
		retH = *retH.Next
	}
	fmt.Println(retH)

}

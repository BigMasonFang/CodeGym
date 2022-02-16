package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	// boundary
	if head == nil || k == 0 || head.Next == nil {
		return head
	}
	tailNode, length := head, 1
	for tailNode.Next != nil {
		tailNode = tailNode.Next
		length++
	}

	if length == 1 {
		return head
	} else if k >= length {
		k = k % length
	}

	// find
	pNode := &ListNode{Val: 0, Next: head}
	for i := 0; i < length-k; i++ {
		pNode = pNode.Next
	}
	tailNode.Next = head
	targetNode := pNode.Next
	pNode.Next = nil

	return targetNode
}

func main() {
	n5 := ListNode{Val: 5, Next: nil}
	n4 := ListNode{Val: 4, Next: &n5}
	n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}

	retH := *rotateRight(&n1, 6)

	for retH.Next != nil {
		fmt.Println(retH)
		retH = *retH.Next
	}
	fmt.Println(retH)
}

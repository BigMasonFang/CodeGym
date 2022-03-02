package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// reverse
	h := reverse(head)
	p := h
	i := 1

	for i < n-1 {
		p = p.Next
		i++
	}
	// skip
	if n == 1 {
		h = p.Next
	} else {
		target := p.Next.Next
		p.Next = target
	}

	// reverse again
	new_h := reverse(h)
	return new_h
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

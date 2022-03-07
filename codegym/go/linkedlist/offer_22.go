package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	// reverse and then get slice and then reverse again
	reverseH := reverseList(head)
	// now get prev as the reverse head
	p := reverseH
	if k == 1 {
		p.Next = nil
		return p
	}
	for k > 1 {
		p = p.Next
		k--
	}
	p.Next = nil
	// reverse again
	trueH := reverseList(reverseH)
	return trueH
}

func reverseList(head *ListNode) *ListNode {
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

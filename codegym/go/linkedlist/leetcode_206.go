package main

func reverseList(head *ListNode) *ListNode {
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

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode1(headA, headB *ListNode) bool {
	// reverse one list and see if they collide
	// actually not collide because the structure is changed
	// so if there is a interesection, the result is
	// u run through the other list, the tail node is the head
	// of the reversed one
	// reverse listA
	reversePA, pB, rst := reverseList(headA), headB, false
	var prevB *ListNode
	for pB != nil {
		prevB = pB
		pB = pB.Next
	}
	if prevB == reversePA {
		rst = true
	}
	reverseList(reversePA)
	return rst
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// border
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	// this condition
	// using pA != pB, why?
	// if intersect, it will end
	// if not intersect, both pA and pB will be nil, it will end too
	for pA != pB {
		if pA == nil {
			pA = headB
			pB = pB.Next
			continue
		}
		if pB == nil {
			pB = headA
			pA = pA.Next
			continue
		}
		pA, pB = pA.Next, pB.Next
	}
	return pA
}

// standard
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func main() {
	c3 := ListNode{Val: 62, Next: nil}
	c2 := ListNode{Val: 61, Next: &c3}
	c1 := ListNode{Val: 60, Next: &c2}

	b3 := ListNode{Val: 72, Next: &c1}
	b2 := ListNode{Val: 71, Next: &b3}
	b1 := ListNode{Val: 70, Next: &b2}

	a2 := ListNode{Val: 100, Next: &c1}
	a1 := ListNode{Val: 99, Next: &a2}

	rstNode := getIntersectionNode(&a1, &b1)
	fmt.Println(rstNode)
}

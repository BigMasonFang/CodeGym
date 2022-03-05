//  Linked List Cycle
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// O(N) o(N)
func hasCycle1(head *ListNode) bool {
	hashMap := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := hashMap[head]; ok {
			return true
		}
		hashMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

// Floyd Cycle Detection Algorithm
// O(N) O(1)
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	p1, p2 := head, head.Next
	for p1 != p2 {
		if p2 == nil || p2.Next == nil {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return true
}

// 20200305 redo
// naive way
func hasCycle_naive(head *ListNode) bool {
	hashMap := make(map[*ListNode]int)
	for head != nil {
		if _, ok := hashMap[head.Next]; ok {
			return true
		} else {
			hashMap[head.Next] = 1
			head = head.Next
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

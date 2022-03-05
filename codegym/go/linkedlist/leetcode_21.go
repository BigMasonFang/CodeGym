package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		return nil
	}
	head := &ListNode{}
	p1, p2, p := list1, list2, head

	// go true && nil?
	for p1.Next != nil && p2.Next != nil {
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
		fmt.Println(p, p1, p2)
	}
	if p1.Next != nil {
		p.Next = p1
	} else if p2.Next != nil {
		p.Next = p2
	}
	return head.Next
}

// 20220305 redo
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// border
	if list1 == nil || list2 == nil {
		if list1 != nil {
			return list1
		}
		if list2 != nil {
			return list2
		}
		return nil
	}

	var rst *ListNode
	if list1.Val <= list2.Val {
		rst = list1
	} else {
		rst = list2
	}
	prev := &ListNode{}
	p1, p2 := list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			prev.Next = p1
			p1 = p1.Next
		} else {
			prev.Next = p2
			p2 = p2.Next
		}
		prev = prev.Next
		fmt.Println(p1, p2, prev)
	}

	if p1 == nil {
		prev.Next = p2
	} else {
		prev.Next = p1
	}

	return rst
}

func main() {
	m3 := ListNode{Val: 3, Next: nil}
	m2 := ListNode{Val: 2, Next: &m3}
	m1 := ListNode{Val: 1, Next: &m2}
	n3 := ListNode{Val: 4, Next: nil}
	n2 := ListNode{Val: 3, Next: &n3}
	n1 := ListNode{Val: 2, Next: &n2}

	retH := mergeTwoLists(&m1, &n1)
	fmt.Println("Result")
	for retH.Next != nil {
		fmt.Println(retH)
		retH = retH.Next
	}
	fmt.Println(retH)

	retH2 := mergeTwoLists(nil, &n1)
	fmt.Println("Result 2")
	for retH.Next != nil {
		fmt.Println(retH2)
		retH2 = retH2.Next
	}
	fmt.Println(retH2)

}

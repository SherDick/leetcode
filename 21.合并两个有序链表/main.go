package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
	时间复杂度：O(n+m)
	空间复杂度：O(1)
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	p, q := l1, l2
	head, pre := &ListNode{}, &ListNode{}
	head.Next = nil
	pre.Next = nil
	if p.Val < q.Val {
		head = p
		p = p.Next
	} else {
		head = q
		q = q.Next
	}
	pre = head
	for p != nil && q != nil {
		if p.Val < q.Val {
			pre.Next = p
			pre = p
			p = p.Next
		} else {
			pre.Next = q
			pre = q
			q = q.Next
		}
	}
	if p == nil {
		pre.Next = q
	} else {
		pre.Next = p
	}
	return head
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  6,
					Next: nil,
				},
			},
		},
	}
	m := mergeTwoLists(l1, l2)
	for m != nil {
		fmt.Printf("%+v ", m.Val)
		m = m.Next
	}

}

package main

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	p, q := l1, l2
	flag := 0
	x, y := 0, 0
	for p != nil || q != nil {
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		num := x + y + flag
		flag = num / 10
		curr.Next = &ListNode{
			Val:  num % 10,
			Next: nil,
		}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if flag > 0 {
		curr.Next = &ListNode{
			Val:  flag,
			Next: nil,
		}
	}
	return head.Next
}

func main() {

}

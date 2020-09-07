package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return judge(root, root)
}

func judge(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return judge(p.Left, q.Right) && judge(p.Right, q.Left)
}

/* 迭代 */
func isSymmetric2(root *TreeNode) bool {
	p, q := root, root
	queue := make([]*TreeNode, 0)
	queue = append(queue, p)
	queue = append(queue, q)

	for len(queue) > 0 {
		p, q = queue[0], queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		queue = append(queue, p.Left)
		queue = append(queue, q.Right)

		queue = append(queue, p.Right)
		queue = append(queue, q.Left)
	}
	return true
}

func main() {

}

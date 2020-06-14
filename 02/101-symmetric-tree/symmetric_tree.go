package binarytree

//https://leetcode-cn.com/problems/symmetric-tree/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirrorIterative(root.Left, root.Right)
}

func isMirrorRecursive(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isMirrorRecursive(left.Left, right.Right) && isMirrorRecursive(left.Right, right.Left)
}

func isMirrorIterative(left, right *TreeNode) bool {
	q := []*TreeNode{}
	q = append(q, left, right)

	for len(q) > 0 {
		n1, n2 := q[0], q[1]
		q = q[2:]

		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}

		q = append(q, n1.Left, n2.Right)
		q = append(q, n1.Right, n2.Left)
	}
	return true
}

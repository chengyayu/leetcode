package binarytree

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return inorderIterative(root)
}

func inorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := append(inorderRecursive(root.Left), root.Val)
	rest = append(rest, inorderRecursive(root.Right)...)
	return rest
}

func inorderIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack, rest := Stack([]*TreeNode{root}), []int{}
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur != nil {
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			stack.Push(cur)
			stack.Push(nil)
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
		} else {
			rest = append(rest, stack.Pop().Val)
		}
	}

	return rest
}

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}

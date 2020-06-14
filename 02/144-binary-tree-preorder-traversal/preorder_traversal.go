package binarytree

//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return preorderIterative(root)
}

func preorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := append([]int{root.Val}, preorderRecursive(root.Left)...)
	rest = append(rest, preorderRecursive(root.Right)...)
	return rest
}

func preorderIterative(root *TreeNode) []int {
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
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
			stack.Push(cur)
			stack.Push(nil)
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

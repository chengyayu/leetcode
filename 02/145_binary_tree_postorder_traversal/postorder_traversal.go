package _145_binary_tree_postorder_traversal

//https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	return postorderIterative(root)
}

func postorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := append(postorderRecursive(root.Left), postorderRecursive(root.Right)...)
	rest = append(rest, root.Val)
	return rest
}

func postorderIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack, rest := Stack([]*TreeNode{root}), []int{}
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur != nil {
			stack.Push(cur)
			stack.Push(nil)
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
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

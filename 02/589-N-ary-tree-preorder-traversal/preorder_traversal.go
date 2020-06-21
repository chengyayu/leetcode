package narytree

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	rest := []int{}
	if root == nil {
		return rest
	}

	rest = append(rest, root.Val)
	for _, c := range root.Children {
		rest = append(rest, preorder(c)...)
	}

	return rest
}

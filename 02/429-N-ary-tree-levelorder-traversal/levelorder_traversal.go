package narytree

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	rest := [][]int{}
	queue := []*Node{root}

	for len(queue) > 0 {
		level := []int{}
		num := len(queue)
		for i := 0; i < num; i++ {
			n := queue[0]
			queue = append(queue[1:], n.Children...)
			level = append(level, n.Val)
		}
		rest = append(rest, level)
	}

	return rest
}

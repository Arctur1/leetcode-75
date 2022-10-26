package binary_search_tree

import "leetcode/tree"

// func lowestCommonAncestor(root, p, q *tree.TreeNode) *tree.TreeNode {
// 	p_anc, q_anc := ancestors(root,p.Val), ancestors(root, q.Val)
	

// }

func ancestors(root *tree.TreeNode, n int) []*tree.TreeNode {
	traversed := []*tree.TreeNode{root}

	for {
		node := traversed[len(traversed)-1]
		if node.Val == n {
			break
		}

		if node.Left != nil && n > node.Val {
			traversed = append(traversed, node.Right)
		} else if node.Right != nil {
			traversed = append(traversed, node.Left)
		}
	}

	return traversed
}

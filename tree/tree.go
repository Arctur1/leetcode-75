package tree

// binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// n-ary tree
type Node struct {
	Val      int
	Children []*Node
}

func create_n_ary_tree(arr []interface{}) *Node {
	root := &Node{arr[0].(int), nil}
	res := root
	for _, value := range arr[2:] {
		if value == nil {
			root = root.Children[0]
			continue
		}
		root.Children = append(root.Children, &Node{value.(int), nil})
	}
	return res
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}
	for _, v := range root.Children {
		tmp := preorder(v)
		if tmp != nil {
			res = append(res[:], tmp...)
		}
	}

	return res
}

func create_binary_tree(arr []interface{}) *TreeNode {
	root := &TreeNode{arr[0].(int), nil, nil}
	nodes := []*TreeNode{root}
	arr = arr[1:]

	for {
		node := nodes[0]
		if arr[0] != nil {
			node.Left = &TreeNode{arr[0].(int), nil, nil}
			nodes = append(nodes, node.Left)
		}

		if arr[1] != nil {
			node.Right = &TreeNode{arr[1].(int), nil, nil}
			nodes = append(nodes, node.Right)
		}

		arr = arr[2:]
		nodes = nodes[1:]

		if len(arr) == 0 || len(nodes) == 0 {
			return root
		}
	}

}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := [][]*TreeNode{}
	nodes = append(nodes, []*TreeNode{root})
	result := [][]int{}
	result = append(result, []int{root.Val})

	level := nodes[0]
	for len(level) > 0 {
		values := []int{}
		next_level := []*TreeNode{}

		for _, node := range level {
			if node.Left != nil {
				next_level = append(next_level, node.Left)
				values = append(values, node.Left.Val)
			}

			if node.Right != nil {
				next_level = append(next_level, node.Right)
				values = append(values, node.Right.Val)
			}
		}
		level = next_level
		if len(values) > 0 {
			result = append(result, values)
		}
	}
	return result
}

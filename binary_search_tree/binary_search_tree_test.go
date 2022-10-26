package binary_search_tree

import (
	"leetcode/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func traverse_list_of_nodes(nodes []*tree.TreeNode) []int {
	result := []int{}
	for _, node := range nodes {
		result = append(result, node.Val)
	}
	return result
}

func TestFindAncestors(t *testing.T) {
	root := tree.Create_binary_tree([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5})
	ancestors_values := traverse_list_of_nodes(ancestors(root, 5))
	assert.Equal(t, ancestors_values, []int{6, 2, 4, 5})
}

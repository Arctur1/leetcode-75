package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorder(t *testing.T) {
	tree := create_n_ary_tree([]interface{}{1, nil, 3, 2, 4, nil, 5, 6})
	assert.Equal(t, preorder(tree), []int{1, 3, 5, 6, 2, 4})
}

func TestLevelOrder(t *testing.T) {
	root := create_binary_tree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	assert.Equal(t, levelOrder(root), [][]int{{3}, {9, 20}, {15, 7}})
}

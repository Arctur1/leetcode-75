package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeLists(t *testing.T) {
	assert := assert.New(t)

	list1 := create_linked_list([]int{1, 2, 4})
	list2 := create_linked_list([]int{1, 3, 4})
	merge := mergeTwoLists(list1, list2)
	// recursive_merge := mergeTwoListsRecursive(list1, list2)

	expected := []int{1, 1, 2, 3, 4, 4}
	// assert.Equal(recursive_merge.traverse_linked_list(), expected, "The two lists should be the same.")
	assert.Equal(merge.traverse_linked_list(), expected, "The two lists should be the same.")

}

func TestDetectCycle(t *testing.T) {
	list := create_linked_list([]int{3, 2, 0, -4})
	list.node(1).Next = list.node(1)
	cycle := detectCycle(list).Val
	assert.Equal(t, cycle, 2, "Cycle starts at value two")
}

func TestReverseLinkedList(t *testing.T) {
	list := create_linked_list([]int{1, 2, 4})
	reversed := reverseList(list)
	assert.Equal(t, reversed.traverse_linked_list(), []int{4, 2, 1})
}

func TestMiddleOfLinkedList(t *testing.T) {
	list := create_linked_list([]int{1, 2, 3, 4, 5})
	middle := middleNode(list)
	assert.Equal(t, middle.traverse_linked_list(), []int{3, 4, 5})

}

package prefix_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningSum(t *testing.T) {
	assert.Equal(t, runningSum([]int{1, 2, 3, 4}), []int{1, 3, 6, 10})
}

func TestPivotIndex(t *testing.T) {
	assert.Equal(t, pivotIndex([]int{1, 7, 3, 6, 5, 6}), 3)
}

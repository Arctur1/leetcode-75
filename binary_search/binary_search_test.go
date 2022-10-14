package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TesBinarySearch(t *testing.T) {
	assert.Equal(t, search([]int{-1, 0, 3, 5, 9, 12}, 9), 4)
}

func TestFirstBadVersion(t *testing.T) {
	isBadVersion = VersionChecker{4}.isBad
	assert.Equal(t, firstBadVersion(5), 4)
}

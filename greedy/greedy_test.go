package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
}

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, longestPalindrome("abccccdd"), 7)

}

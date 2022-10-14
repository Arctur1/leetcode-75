package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIsomorphic(t *testing.T) {
	assert.Equal(t, isIsomorphic("egg", "add"), true)
}

func TestIsSubsequence(t *testing.T) {
	assert.Equal(t, isSubsequence("abc", "ahbgdc"), true)

}

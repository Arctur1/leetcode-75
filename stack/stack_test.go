package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackSpaceCompare(t *testing.T) {
	assert.Equal(t, true, backspaceCompare("ab#c", "ad#c"))
	assert.Equal(t, true, backspaceCompare("ab##", "c#d#"))
	assert.Equal(t, true, backspaceCompare("a##c", "#a#c"))
}

func TestDecodeString(t *testing.T) {
	assert.Equal(t, "aaabcbc", decodeString("3[a]2[bc]"))
}

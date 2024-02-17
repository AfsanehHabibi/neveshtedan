package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenAnArrayWithNilElementsIsGivenRemoveNilRemovesTheNilElements(t *testing.T) {
	arr := []*int{new(int), new(int), nil, new(int)}
	assert.Len(t, RemoveNilElements[int](arr), 3)
}

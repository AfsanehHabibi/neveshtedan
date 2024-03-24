package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenAnArrayWithNilElementsIsGivenRemoveNilRemovesTheNilElements(t *testing.T) {
	var a = 1
	var b = 2
	var c = 3
	arr := []*int{&a, &b, nil, &c}
	result := RemoveNilElements(arr)
	assert.Len(t, result, 3)
	for i := 0; i < 3; i++ {
		assert.Equal(t, i+1, result[i])
	}
}

func TestWhenAnArrayIsGivenConvertToPointerPreserveValues(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := ConvertToPointerArray(arr)
	assert.Len(t, result, 4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, i+1, *result[i])
	}
}

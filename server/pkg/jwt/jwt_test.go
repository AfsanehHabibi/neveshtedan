package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTokenGeneratesTokenWithoutError(t *testing.T) {
	id := 123
	token, err := GenerateToken(id)
	assert.NoError(t, err)
	outId, err := ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, id, outId)
}

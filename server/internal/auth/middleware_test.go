package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromContextReturnUserIdCorrectly(t *testing.T) {
	id := 35983
	ctx := AddUserToContext(context.Background(), id)
	outId := GetUseFromContext(ctx)
	assert.Equal(t, id, *outId)
}

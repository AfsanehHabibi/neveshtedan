package repository_test

import (
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestWhenAddingMultipleUserCanRetrieveItLater(t *testing.T) {
	id, err := userRep.Add(ctx, model.NewUser{Username: "ali", Password: "eo9h49g84"})
	assert.NoError(t, err)

	user, err := userRep.GetById(ctx, id)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "ali", user.Username)
	userRep.Clear(ctx)
}

func TestWhenAddingTwoUserTheirIdIsNotEqual(t *testing.T) {
	id1, err := userRep.Add(ctx, model.NewUser{Username: "ali", Password: "eo9h49g84"})
	assert.NoError(t, err)
	id2, err := userRep.Add(ctx, model.NewUser{Username: "mohammad", Password: "eo9h489g84"})
	assert.NoError(t, err)

	assert.NotEqual(t, id1, id2)
	userRep.Clear(ctx)
}

func TestWhenAddingAUserWithDuplicateUsernameShouldReceiveError(t *testing.T) {
	_, err := userRep.Add(ctx, model.NewUser{Username: "ali", Password: "eo9h49g84"})
	assert.NoError(t, err)

	_, err = userRep.Add(ctx, model.NewUser{Username: "ali", Password: "kf9rghr8j"})
	assert.Error(t, err)
	userRep.Clear(ctx)
}

func TestWhenExistentUsernameAndPasswordIsGivenCanGetId(t *testing.T) {
	id, err := userRep.Add(ctx, model.NewUser{Username: "ali", Password: "eo9h49g84"})
	assert.NoError(t, err)

	inputId, err := userRep.GetIdIfExists(ctx, "ali", "eo9h49g84")
	assert.NoError(t, err)
	assert.NotNil(t, inputId)
	assert.Equal(t, id, *inputId)
	userRep.Clear(ctx)
}

func TestWhenWrongPasswordIsGivenCanNotGetId(t *testing.T) {
	_, err := userRep.Add(ctx, model.NewUser{Username: "ali", Password: "eo9h49g84"})
	assert.NoError(t, err)

	inputId, err := userRep.GetIdIfExists(ctx, "ali", "wrongpass")
	assert.NoError(t, err)
	assert.Nil(t, inputId)
	userRep.Clear(ctx)
}

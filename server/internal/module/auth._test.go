package module

import (
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestWhenNewUserCreatesItReceivesToken(t *testing.T) {
	defer clear()

	token, err := module.CreateUser(ctx, model.NewUser{Username: "Ahmad", Password: "03jf9efk"})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestWhenExistentUserLogsInItReceivesToken(t *testing.T) {
	defer clear()

	username := "Ahmad"
	password := "03jf9efk"
	token, err := module.CreateUser(ctx, model.NewUser{Username: username, Password: password})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	tokenL, err := module.Login(ctx, model.Login{Username: username, Password: password})
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenL)
}

func TestWhenWrongPasswordEnteredLoginReturnsError(t *testing.T) {
	defer clear()

	username := "Ahmad"
	password := "03jf9efk"
	token, err := module.CreateUser(ctx, model.NewUser{Username: username, Password: password})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	tokenL, err := module.Login(ctx, model.Login{Username: username, Password: "wrong"})
	assert.Error(t, err)
	assert.Empty(t, tokenL)
}
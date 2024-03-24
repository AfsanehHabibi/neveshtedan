package module

import (
	"context"
	"errors"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/jwt"
)

func (m NeveshtedanModule) Login(ctx context.Context, input model.Login) (string, error) {
	id, err := m.userRep.GetIdIfExists(ctx, input.Username, input.Password)
	if err != nil {
		return "", err
	}
	if id == nil {
		return "", errors.New("wrong password or username")
	}
	token, err := jwt.GenerateToken(*id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (m NeveshtedanModule) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	id, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", errors.New("access denied")
	}
	token, err := jwt.GenerateToken(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

package module

import (
	"context"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/jwt"
)

func (m NeveshtedanModule) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	id, err := m.userRep.Add(ctx, input)
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(id)
	if err != nil {
		return "", err
	}
	return token, nil
}

package repository

import (
	"context"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
)

type UserRepository interface {
	Add(ctx context.Context, user model.NewUser) (int, error)
	GetById(ctx context.Context, id int) (*model.User, error)
	Clear(ctx context.Context) error
}
